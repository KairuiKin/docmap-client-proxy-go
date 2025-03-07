//go:build windows

package tunnel

import (
	"context"
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"encoding/binary"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sys/windows"
)

type windowsPipeServer struct {
	PipeServerBase
	pipeName      string
	handles       map[string]windows.Handle
	mu            sync.RWMutex
	clientCounter uint64
	wg            sync.WaitGroup
	ctx           context.Context
	cancel        context.CancelFunc
}

func NewPipeServer(pipeName string) PipeServer {
	ctx, cancel := context.WithCancel(context.Background())
	if !strings.HasPrefix(pipeName, `\\.\pipe\`) {
		pipeName = `\\.\pipe\` + pipeName
	}
	return &windowsPipeServer{
		pipeName: pipeName,
		handles:  make(map[string]windows.Handle),
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (ps *windowsPipeServer) Start() error {
	err := ps.PipeServerBase.Start()
	if err != nil {
		return err
	}

	ps.mu.Lock()
	ps.running = true
	ps.mu.Unlock()

	numCPU := runtime.NumCPU()
	//开启多个监听
	for i := 0; i < numCPU; i++ {
		ps.wg.Add(1)
		go func() {
			defer ps.wg.Done()
			ps.listen()
		}()
	}
	return nil
}

func (ps *windowsPipeServer) listen() {
	for ps.running {
		select {
		case <-ps.ctx.Done():
			return
		default:
			pipe, err := windows.CreateNamedPipe(
				windows.StringToUTF16Ptr(ps.pipeName),
				windows.PIPE_ACCESS_DUPLEX,
				windows.PIPE_TYPE_MESSAGE|windows.PIPE_WAIT,
				windows.PIPE_UNLIMITED_INSTANCES,
				1024,
				1024,
				0,
				nil,
			)
			if err != nil {
				common.Logger.Infof("Failed to create named pipe:", err)
				continue // 继续尝试创建新的管道
			}
			connected := false
			if err := windows.ConnectNamedPipe(pipe, nil); err != nil {
				if !errors.Is(err, windows.ERROR_PIPE_CONNECTED) {
					fmt.Println("Failed to connect named pipe:", err)
					windows.CloseHandle(pipe)
					continue
				}
				connected = true
			} else {
				// ConnectNamedPipe 成功连接
				connected = true
			}

			if connected {
				ps.mu.Lock()
				ps.clientCounter++
				clientID := fmt.Sprintf("client-%d", ps.clientCounter)
				ps.handles[clientID] = pipe
				ps.mu.Unlock()

				ps.handleClient(clientID, pipe)
			}
		}

	}
}

func (ps *windowsPipeServer) handleClient(clientID string, pipe windows.Handle) {
	defer func() {
		ps.mu.Lock()
		delete(ps.handles, clientID)
		ps.mu.Unlock()
		err := windows.CloseHandle(pipe)
		if err != nil {
			fmt.Printf("Failed to close pipe for client %s: %v\n", clientID, err)
		}
	}()

	for ps.running {
		// 读取消息头（假定消息头固定为 4 字节，包含消息体长度等信息）
		header := make([]byte, 4)
		var bytesRead uint32
		err := windows.ReadFile(pipe, header, &bytesRead, nil)
		if err != nil {
			if errors.Is(err, windows.ERROR_BROKEN_PIPE) {
				fmt.Printf("Client %s disconnected\n", clientID)
			} else {
				fmt.Printf("Error reading header from client %s: %v\n", clientID, err)
			}
			break
		}

		// 确保读取到完整的头部
		if bytesRead < 4 {
			fmt.Printf("Incomplete header received from client %s\n", clientID)
			break
		}

		bodyLength, err := DecodeHeader(header[:4])
		// 解析消息体长度（假设长度存储在前 4 字节，Little Endian）
		if bodyLength == 0 {
			fmt.Printf("Invalid body length from client %s\n", clientID)
			break
		}

		// 分配缓冲区并读取消息体
		body := make([]byte, bodyLength)
		err = windows.ReadFile(pipe, body, &bytesRead, nil)
		if err != nil {
			fmt.Printf("Error reading body from client %s: %v\n", clientID, err)
			break
		}

		// 确保读取到完整的消息体
		if bytesRead < bodyLength {
			fmt.Printf("Incomplete body received from client %s\n", clientID)
			break
		}

		// 通知观察者收到完整消息
		ps.NotifyObservers(clientID, body)
	}
}

func (ps *windowsPipeServer) Send(clientID string, message []byte) error {
	ps.mu.RLock()
	pipe, exists := ps.handles[clientID]
	ps.mu.RUnlock()
	if !exists {
		return errors.New("client not found")
	}

	// Calculate the length of the message
	messageLength := uint32(len(message))

	// Create a buffer to hold the length and the message
	buffer := make([]byte, 4+len(message))

	// Encode the message length in little endian format
	binary.BigEndian.PutUint32(buffer[:4], messageLength)

	// Copy the message into the buffer after the length
	copy(buffer[4:], message)

	var bytesWritten uint32

	// Write the combined buffer to the pipe in a single operation
	err := windows.WriteFile(pipe, buffer, &bytesWritten, nil)
	if err != nil {
		return err
	}

	// Optional: Verify that all bytes were written
	if bytesWritten != uint32(len(buffer)) {
		return errors.New("incomplete write to pipe")
	}

	return nil
}

func (ps *windowsPipeServer) Stop() error {
	err := ps.PipeServerBase.Stop()
	if err != nil {
		return err
	}

	ps.cancel() // 发送取消信号

	ps.mu.Lock()
	ps.running = false
	// 不再关闭句柄，handleClient 会处理
	ps.mu.Unlock()

	// 等待所有 goroutine 退出
	ps.wg.Wait()
	return nil
}

// NotifyObservers 通知所有回调
func (ps *windowsPipeServer) NotifyObservers(clientID string, message []byte) {
	var parsy model.ExchangeMessage
	err := proto.Unmarshal(message, &parsy)
	if err != nil {
		common.Logger.Errorf("Received message faild %v", err)
		return
	}
	// 使用 strconv.ParseUint 将字符串解析为 uint64
	uint32Value, err := strconv.ParseUint(parsy.Me.ProcessId, 10, 64)
	if err != nil {
		fmt.Printf("转换错误: %v\n", err)
		return
	}
	client := common.GetClientInstance()
	client.AddClient(parsy.Me.ClientInstanceId, parsy.Me, uint(uint32Value))
	if !parsy.Response {
		switch parsy.Type {
		case model.ExchangeMessage_EVENT_ACCESS_PERMISSION:
			response := &model.EventAccessPermissionResponse{
				Access:   model.EventAccessPermissionResponse_allow,
				PolicyId: 0,
				Result:   true,
			}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		case model.ExchangeMessage_REPORT_ACTION_EVENT:
			response := &model.ReportEventResponse{
				Result: true,
			}

			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
			//事件转移至event包处理
			// 发送异步消息
			var msg *common.Message
			msg = &common.Message{
				SourceID:    "tunnel",
				TargetID:    "event",
				Destination: common.MessageLog,
				Type:        common.MessageInfo,
				Payload:     message,
			}
			bus := messagebus.GetMessageBusInstance()
			bus.SubmitAsyncTask(msg)
		case model.ExchangeMessage_PUT_TAG:
			response := &model.PutTagResponse{}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		case model.ExchangeMessage_GET_TAG:
			response := &model.GetTagResponse{}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		case model.ExchangeMessage_REDIRECT:
			response := &model.RedirectResponse{}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		case model.ExchangeMessage_DECRYPT:
			response := &model.DecryptResponse{}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		case model.ExchangeMessage_GET_OUTDOOR_PATH:
			response := &model.GetOutDoorPathResponse{}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		case model.ExchangeMessage_REFRESH_PAUSE:
			response := &model.RefreshPauseResponse{}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		case model.ExchangeMessage_SUBMIT_USERNAMES:
			response := &model.SubmitUserNamesResponse{
				Result: true,
			}
			if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
				return
			}
		}
	} else {
		fmt.Printf("Received message from %s:type%v %s\n", clientID, parsy.Type, string(message))
	}
	//
}

func (ps *windowsPipeServer) createAndSendResponse(clientID string, parsy *model.ExchangeMessage, payload proto.Message) error {
	anyPayload, err := anypb.New(payload)
	if err != nil {
		common.Logger.Errorf("Failed to marshal payload: %v", err)
		return err
	}

	responseMessage := &model.ExchangeMessage{
		Id:       parsy.Id,
		Uuid:     parsy.GetUuid(),
		Type:     parsy.GetType(),
		Response: true,
		Payload:  anyPayload,
	}

	messagePayload, err := proto.Marshal(responseMessage)
	if err != nil {
		common.Logger.Errorf("Failed to marshal response message: %v", err)
		return err
	}

	return ps.Send(clientID, messagePayload)
}
