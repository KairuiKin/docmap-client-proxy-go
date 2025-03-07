package tunnel

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/proto/docmap/message/model"
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
	"strconv"
	"time"
)

const (
	headerLength = 4 // 消息头长度，表示消息体的长度
)

// DMApiMessageType 定义了消息类型的枚举值
type DMApiMessageType int

const (
	// Heartbeat 心跳消息
	Heartbeat DMApiMessageType = 0x1

	// Whoami 客户端上报身份
	Whoami DMApiMessageType = 0x2
	// Whoareu 服务器询问客户端身份
	Whoareu DMApiMessageType = 0x3

	// MarkDocSecret 标记文档密级
	MarkDocSecret DMApiMessageType = 0x4
	// MarkDocSecretResult 标记文档密级执行结果
	MarkDocSecretResult DMApiMessageType = 0x5
	// MarkDocClassification 标记文档类别
	MarkDocClassification DMApiMessageType = 0x6
	// MarkDocClassificationResult 标记文档类别执行结果
	MarkDocClassificationResult DMApiMessageType = 0x7

	// MarkFamilySecret 批量标记文档家族密级
	MarkFamilySecret DMApiMessageType = 0x8
	// MarkFamilySecretProgress 批量标记文档家族密级执行进度
	MarkFamilySecretProgress DMApiMessageType = 0x9
	// MarkFamilyClassification 批量标记文档家族类别
	MarkFamilyClassification DMApiMessageType = 0xa
	// MarkFamilyClassificationProgress 批量标记文档家族类别执行进度
	MarkFamilyClassificationProgress DMApiMessageType = 0xb

	// SubmitMarkFamilySecret 提交批量标记文档家族密级请求
	SubmitMarkFamilySecret DMApiMessageType = 0xc
	// SubmitMarkFamilyClassification 提交批量标记文档家族类别请求
	SubmitMarkFamilyClassification DMApiMessageType = 0xd

	// ListClientFileInfo 服务端查询客户端指定目录下的文件和子目录清单
	ListClientFileInfo DMApiMessageType = 0xe
	// ListClientFileInfoResult 客户端上报指定目录下的文件和子目录清单
	ListClientFileInfoResult DMApiMessageType = 0xf

	// MarkDirectoryClassification 服务端标记目录类别
	MarkDirectoryClassification DMApiMessageType = 0x10
	// MarkDirectoryClassificationResult 客户端上报标记目录类别执行结果
	MarkDirectoryClassificationResult DMApiMessageType = 0x11
	// MarkDirectorySecret 服务端标记目录密级
	MarkDirectorySecret DMApiMessageType = 0x12
	// MarkDirectorySecretResult 客户端上报标记目录密级执行结果
	MarkDirectorySecretResult DMApiMessageType = 0x13

	// UninstallClient 服务端卸载客户端
	UninstallClient DMApiMessageType = 0x14
	// SendClientIPPolicy 服务端发送知识产权列表
	SendClientIPPolicy DMApiMessageType = 0x15
	// SendClientOutdoorKeyInfo 服务端发送外发密钥列表
	SendClientOutdoorKeyInfo DMApiMessageType = 0x16
	// SendClientDecryptPassword 服务端发送加密文件解密密码
	SendClientDecryptPassword DMApiMessageType = 0x17
	// SendClientConfigInfo 服务端发送配置信息
	SendClientConfigInfo DMApiMessageType = 0x18
	// SendClientConfigInfoInner 服务端发送办公软件规则配置信息
	SendClientConfigInfoInner DMApiMessageType = 0x19
	// SendClientConfigInfoOnline 服务端发送需要支持的临时文档后缀清单
	SendClientConfigInfoOnline DMApiMessageType = 0x1A
)

type TCPHandler struct {
	address           string
	Timeout           time.Duration
	retryInterval     time.Duration
	heartbeatInterval time.Duration
	seedId            uint64
	Conn              net.Conn // 客户端连接
}

// NewTCPHandler 创建TCP客户端处理实例
func NewTCPHandler(address string) *TCPHandler {
	return &TCPHandler{
		address:           address,
		Timeout:           time.Second * 10,
		retryInterval:     3 * time.Second,  // 重连间隔
		heartbeatInterval: 10 * time.Second, // 心跳包发送间隔
		seedId:            0,                // 心跳包内容，可根据需要修改

	}
}

// Start 尝试连接服务端并在连接断开后自动重连
func (t *TCPHandler) Start(messageCh chan<- *common.Message) {
	go func() {
		for {
			common.Logger.Infof("尝试连接 %s", t.address)
			conn, err := net.Dial("tcp", t.address)
			if err != nil {
				common.Logger.Warnf("连接服务器失败: %v. %s 后重新连接 ", err, t.retryInterval)
				time.Sleep(t.retryInterval)
				continue
			}
			common.Logger.Infof("连接服务器: %s", t.address)
			t.Conn = conn
			// 启动读取与心跳协程
			done := make(chan struct{})
			go t.startReading(conn, messageCh, done)
			go t.startHeartbeat(done)

			// 阻塞等待读取协程结束（表示连接断开）
			<-done
			common.Logger.Warnf("连接断开, 即将重新连接...")
			// 重连逻辑在for循环中继续执行
		}
	}()
}

// startReading 从服务器持续读取数据，一旦出现错误表示断线，则通知上层进行重连
func (t *TCPHandler) startReading(conn net.Conn, messageCh chan<- *common.Message, done chan<- struct{}) {
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
		close(done)
	}()
	for {
		header := make([]byte, headerLength)
		var headerRead int // 记录已读取的字节数
		for headerRead < int(headerLength) {
			n, err := conn.Read(header[headerRead:]) // 从上次读取的位置继续读取
			if err != nil {
				return
			}
			headerRead += n // 更新已读取的字节数
		}

		// 解析消息体长度
		bodyLength := binary.BigEndian.Uint32(header)

		// 读取消息体
		body := make([]byte, bodyLength)
		var totalRead int // 记录已读取的字节数
		for totalRead < int(bodyLength) {
			n, err := conn.Read(body[totalRead:]) // 从上次读取的位置继续读取
			if err != nil {
				return
			}
			totalRead += n // 更新已读取的字节数
		}

		if totalRead == int(bodyLength) {
			err := t.MessagePaser(body)
			if err != nil {
				return
			}
		}
	}
}

// startHeartbeat 定期向服务器发送心跳包，当发送失败表示连接已断开
func (t *TCPHandler) startHeartbeat(done <-chan struct{}) {
	ticker := time.NewTicker(t.heartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			// 当阅读协程结束(done关闭)，连接也已经断开，退出心跳循环
			return
		case <-ticker.C:
			client := common.GetClientInstance()
			var sendMessage model.DMExchangeMessage
			sendMessage.Id = t.GetMessageId()
			sendMessage.Type = uint32(Heartbeat)
			sendMessage.Uuid, _ = common.DMRandomUuid()
			sendMessage.Devid = &client.DeviceID
			sendMessage.Response = false
			serializedData, err := proto.Marshal(&sendMessage)
			err = t.SendPayload(serializedData)
			if err != nil {
				common.Logger.Warnf("Failed to send heartbeat: %v", err)
				// 写入失败意味着连接问题，结束心跳协程，让startReading结束通知重连
				return
			} else {
				//logger.Info("Heartbeat sent: %s", serializedData)
			}
		}
	}
}

func (t *TCPHandler) GetMessageId() uint64 {
	t.seedId++
	return t.seedId
}

// SendToken 发送 token 到服务器
func (t *TCPHandler) SendPayload(body []byte) error {
	if t.Conn == nil {
		common.Logger.Errorf("Failed to connect to %s", t.address)
		return fmt.Errorf("未建立连接，请先调用 Start 方法")
	}

	// 设置写超时
	err := t.Conn.SetWriteDeadline(time.Now().Add(t.Timeout))
	if err != nil {
		return err
	}

	// 计算消息体长度并创建消息头，固定为 4 个字节
	headLen := 4
	header := make([]byte, headLen)
	binary.BigEndian.PutUint32(header, uint32(len(body)))

	// 创建数据缓冲区，将消息头和消息体合并
	message := append(header, body...)

	// 发送消息
	_, err = t.Conn.Write(message)
	if err != nil {
		return err
	}

	return nil
}

// SendToken 发送 token 到服务器
func (t *TCPHandler) RegisterCallDidline() error {
	if t.Conn == nil {
		common.Logger.Errorf("Failed to connect to %s", t.address)
		return fmt.Errorf("未建立连接，请先调用 Start 方法")
	}

	// 设置写超时
	err := t.Conn.SetWriteDeadline(time.Now().Add(t.Timeout))
	if err != nil {
		return err
	}
	client := common.GetClientInstance()
	var whoAmI model.DMWhoAmIReq
	whoAmI.UserId, _ = strconv.ParseUint(client.UserId, 10, 64)
	whoAmI.AccessToken = client.AccessToken
	whoAmI.DeviceId = client.DeviceID
	var sendMessage model.DMExchangeMessage
	sendMessage.Id = client.GetMessageId()
	sendMessage.Type = uint32(Whoami)
	sendMessage.Uuid, _ = common.DMRandomUuid()
	sendMessage.Devid = &client.DeviceID
	sendMessage.Response = false
	// 序列化 req
	payload, err := anypb.New(&whoAmI)
	if err != nil {
		fmt.Println("序列化 req 失败:", err)
		return err
	}

	sendMessage.Payload = payload
	serializedData, err := proto.Marshal(&sendMessage)
	err = t.SendPayload(serializedData)
	if err != nil {
		return err
	}
	return nil
}

// SendToken 发送 token 到服务器
func (t *TCPHandler) MessagePaser(body []byte) error {
	var message model.DMExchangeMessage
	err := proto.Unmarshal(body, &message)
	if err != nil {
		return err
	}
	switch message.Type {
	case uint32(MarkDocSecret):
	case uint32(MarkDocClassification):
	case uint32(MarkFamilySecret):
	case uint32(MarkFamilyClassification):
	case uint32(ListClientFileInfo):
	case uint32(SendClientIPPolicy):
	case uint32(SendClientConfigInfo):
	case uint32(SendClientConfigInfoInner):
	case uint32(SendClientConfigInfoOnline):
	}
	return nil
}
