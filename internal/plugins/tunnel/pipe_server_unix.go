//go:build !windows

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
	"net"
	"os"
	"strconv"
	"sync"
	"syscall"
)

const (
	MaxMessageSize = 1 << 20 // 1MB
)

// unixPipeServer 是 Unix 平台的管道服务器实现
type unixPipeServer struct {
	PipeServerBase
	socketPath     string
	listener       net.Listener
	clients        sync.Map // 使用 sync.Map 代替 map + mutex
	clientCounter  uint64
	clientCounterM sync.Mutex // 用于保护 clientCounter 的递增
	wg             sync.WaitGroup
	ctx            context.Context
	cancel         context.CancelFunc
}

// NewPipeServer 创建一个新的 Unix PipeServer
func NewPipeServer(socketPath string) PipeServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &unixPipeServer{
		socketPath: socketPath,
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Start 启动 Unix PipeServer
func (ps *unixPipeServer) Start() error {
	err := ps.PipeServerBase.Start()
	if err != nil {
		return err
	}

	// 删除旧的套接字文件（如果存在）
	if _, err := os.Stat(ps.socketPath); err == nil {
		if err := os.Remove(ps.socketPath); err != nil {
			return fmt.Errorf("无法删除旧的套接字文件: %v", err)
		}
	}

	listener, err := net.Listen("unix", ps.socketPath)
	if err != nil {
		return err
	}

	// 设置套接字文件的权限为 0666（所有用户都可以读写）
	if err := syscall.Chmod(ps.socketPath, 0666); err != nil {
		common.Logger.Errorf("Error setting socket file permissions: %v", err)
	}

	common.Logger.Println("Socket created and listening on", ps.socketPath)

	ps.listener = listener
	ps.wg.Add(1)
	go func() {
		defer ps.wg.Done()
		ps.listen()
	}()
	common.Logger.Printf("Unix PipeServer started on %s", ps.socketPath)
	return nil
}

// listen 监听客户端连接
func (ps *unixPipeServer) listen() {
	for {
		select {
		case <-ps.ctx.Done():
			return
		default:
			conn, err := ps.listener.Accept()
			if err != nil {
				// 如果服务器正在关闭，Accept 会返回错误
				select {
				case <-ps.ctx.Done():
					return
				default:
					common.Logger.Errorf("Failed to accept connection: %v", err)
					continue
				}
			}

			// 生成唯一的 clientID
			ps.clientCounterM.Lock()
			ps.clientCounter++
			clientID := fmt.Sprintf("client-%d", ps.clientCounter)
			ps.clientCounterM.Unlock()

			ps.clients.Store(clientID, conn)
			common.Logger.Infof("Accepted connection from client %s", clientID)

			ps.wg.Add(1)
			go func() {
				defer ps.wg.Done()
				ps.handleClient(clientID, conn)
			}()
		}
	}
}

// handleClient 处理客户端连接
func (ps *unixPipeServer) handleClient(clientID string, conn net.Conn) {
	defer func() {
		ps.clients.Delete(clientID)
		err := conn.Close()
		if err != nil {
			common.Logger.Errorf("Failed to close connection for client %s: %v", clientID, err)
		}
		common.Logger.Infof("Connection closed for client %s", clientID)
	}()

	for {
		select {
		case <-ps.ctx.Done():
			return
		default:
			// 读取消息头
			header := make([]byte, 4)
			bytesRead, err := readFull(conn, header)
			if err != nil {
				if errors.Is(err, net.ErrClosed) || errors.Is(err, os.ErrClosed) {
					common.Logger.Errorf("Client %s disconnected", clientID)
				} else {
					common.Logger.Infof("Error reading header from client %s: %v", clientID, err)
				}
				return
			}

			if bytesRead < 4 {
				common.Logger.Errorf("Incomplete header received from client %s", clientID)
				return
			}

			bodyLength, err := DecodeHeader(header[:4])
			if err != nil {
				common.Logger.Errorf("Failed to decode header from client %s: %v", clientID, err)
				return
			}

			if bodyLength == 0 || bodyLength > MaxMessageSize {
				common.Logger.Errorf("Invalid or too large body length (%d) from client %s", bodyLength, clientID)
				return
			}

			// 读取消息体
			body := make([]byte, bodyLength)
			bytesRead, err = readFull(conn, body)
			if err != nil {
				common.Logger.Errorf("Error reading body from client %s: %v", clientID, err)
				return
			}

			if uint32(bytesRead) != bodyLength {
				common.Logger.Errorf("Incomplete body received from client %s", clientID)
				return
			}
			common.Logger.Debugf("Received message from client %s: %s", clientID, string(body))
			// 通知观察者收到完整消息
			ps.NotifyObservers(clientID, body[:bytesRead])
			return
		}
	}
}

// readFull 确保从连接中读取到指定长度的数据
func readFull(conn net.Conn, buf []byte) (int, error) {
	total := 0
	for total < len(buf) {
		n, err := conn.Read(buf[total:])
		if err != nil {
			return total, err
		}
		if n == 0 {
			break
		}
		total += n
	}
	return total, nil
}

// Send 发送消息给指定客户端
func (ps *unixPipeServer) Send(clientID string, message []byte) error {
	value, exists := ps.clients.Load(clientID)
	if !exists {
		return errors.New("client not found")
	}
	conn, ok := value.(net.Conn)
	if !ok {
		return errors.New("invalid connection type")
	}

	// 计算消息长度
	messageLength := uint32(len(message))

	// 创建缓冲区并添加长度前缀
	buffer := make([]byte, 4+len(message))
	binary.BigEndian.PutUint32(buffer[:4], messageLength)
	copy(buffer[4:], message)

	// 发送完整的消息
	_, err := conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

// Stop 停止 Unix PipeServer
func (ps *unixPipeServer) Stop() error {
	err := ps.PipeServerBase.Stop()
	if err != nil {
		return err
	}

	// 取消上下文，通知所有 goroutine 退出
	ps.cancel()

	// 关闭监听器
	if ps.listener != nil {
		err = ps.listener.Close()
		if err != nil {
			common.Logger.Errorf("Failed to close listener: %v", err)
			return err
		}
	}

	// 关闭所有客户端连接
	ps.clients.Range(func(key, value interface{}) bool {
		conn, ok := value.(net.Conn)
		if ok {
			err := conn.Close()
			if err != nil {
				common.Logger.Errorf("Failed to close connection for client %s: %v", key, err)
			}
		}
		ps.clients.Delete(key)
		return true
	})

	// 等待所有 goroutine 退出
	ps.wg.Wait()
	common.Logger.Errorf("Unix PipeServer stopped")
	return nil
}

// NotifyObservers 通知所有回调
func (ps *unixPipeServer) NotifyObservers(clientID string, message []byte) {
	var parsy model.ExchangeMessage
	err := proto.Unmarshal(message, &parsy)
	if err != nil {
		common.Logger.Errorf("Failed to unmarshal message from client %s: %v", clientID, err)
		return
	}

	// 使用 strconv.ParseUint 将字符串解析为 uint64
	uint32Value, err := strconv.ParseUint(parsy.Me.ProcessId, 10, 64)
	if err != nil {
		common.Logger.Errorf("Failed to parse ProcessId for client %s: %v", clientID, err)
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
				common.Logger.Errorf("Failed to send response to client %s: %v", clientID, err)
				return
			}

		case model.ExchangeMessage_REPORT_ACTION_EVENT:
			//response := &model.ReportEventResponse{
			//	Result: true,
			//}
			//if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
			//	common.Logger.Errorf("Failed to send response to client %s: %v", clientID, err)
			//	return
			//}
			// 事件转移至 event 包处理
			// 发送异步消息
			msg := &common.Message{
				SourceID:    "tunnel",
				TargetID:    "event",
				Destination: common.MessageLog,
				Type:        common.MessageInfo,
				Payload:     message,
			}
			bus := messagebus.GetMessageBusInstance()
			bus.SubmitAsyncTask(msg)
		case model.ExchangeMessage_PUT_TAG, model.ExchangeMessage_GET_TAG,
			model.ExchangeMessage_REDIRECT, model.ExchangeMessage_DECRYPT,
			model.ExchangeMessage_GET_OUTDOOR_PATH, model.ExchangeMessage_REFRESH_PAUSE,
			model.ExchangeMessage_SUBMIT_USERNAMES:

			var response proto.Message

			switch parsy.Type {
			case model.ExchangeMessage_PUT_TAG:
				response = &model.PutTagResponse{}
			case model.ExchangeMessage_GET_TAG:
				response = &model.GetTagResponse{}
			case model.ExchangeMessage_REDIRECT:
				response = &model.RedirectResponse{}
			case model.ExchangeMessage_DECRYPT:
				response = &model.DecryptResponse{}
			case model.ExchangeMessage_GET_OUTDOOR_PATH:
				response = &model.GetOutDoorPathResponse{}
			case model.ExchangeMessage_REFRESH_PAUSE:
				response = &model.RefreshPauseResponse{}
			case model.ExchangeMessage_SUBMIT_USERNAMES:
				response = &model.SubmitUserNamesResponse{Result: true}
			}

			if response != nil {
				if err := ps.createAndSendResponse(clientID, &parsy, response); err != nil {
					common.Logger.Errorf("Failed to send response to client %s: %v", clientID, err)
					return
				}
			}

		default:
			common.Logger.Warnf("Unknown message type %v from client %s", parsy.Type, clientID)
		}
	} else {
		common.Logger.Infof("Received message from %s: type %v, %s", clientID, parsy.Type, string(message))
	}
}

// createAndSendResponse 创建并发送响应消息
func (ps *unixPipeServer) createAndSendResponse(clientID string, parsy *model.ExchangeMessage, payload proto.Message) error {
	anyPayload, err := anypb.New(payload)
	if err != nil {
		common.Logger.Errorf("Failed to marshal payload for client %s: %v", clientID, err)
		return err
	}

	responseMessage := &model.ExchangeMessage{
		Id:       parsy.Id, // 使用请求的 Id
		Uuid:     parsy.GetUuid(),
		Type:     parsy.GetType(),
		Response: true,
		Payload:  anyPayload,
	}

	messagePayload, err := proto.Marshal(responseMessage)
	if err != nil {
		common.Logger.Errorf("Failed to marshal response message for client %s: %v", clientID, err)
		return err
	}

	err = ps.Send(clientID, messagePayload)
	if err != nil {
		common.Logger.Errorf("Failed to send response to client %s: %v", clientID, err)
		return err
	}

	return nil
}
