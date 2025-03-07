package tunnel

import (
	"docmap-client-proxy-go/internal/common"
	"fmt"
	"runtime"
)

type CommunicationPlugin struct {
	id          string
	messageCh   chan *common.Message
	commonCh    chan *common.Message
	httpHandler *HTTPHandler
	tcpHandler  *TCPHandler
	pipServer   PipeServer
	httpServer  *HTTPServer
}

func (p *CommunicationPlugin) ID() string {
	return p.id
}

func (p *CommunicationPlugin) Stop() error {
	//TODO implement me
	err := p.pipServer.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (p *CommunicationPlugin) HandleSyncMessage(input *common.Message) (*common.Message, error) {
	var err error
	switch input.Destination {
	case common.MessageLogin:
		// 发送到 TCP 目标
		err = p.Login(input)
	case common.MessageAnonymousLogin:
		// 发送到 TCP 目标
		_, err = p.AnonymousLogin(input)
		if err == nil {
			err = p.tcpHandler.RegisterCallDidline()
		}
	case common.MessageLogout:
		// 发送到 TCP 目标
		err = p.Logout(input)
	case common.MessageStatOuterAudit:
		// 发送到 TCP 目标
		err = p.StatOuterAudit(input)
	case common.MessageOuterAudit:
		// 发送到 TCP 目标
		err = p.OuterAudit(input)
	case common.MessageGetConfig:
		// 发送到 TCP 目标
		err = p.GetConfig(input)
	case common.MessageSendLog:
		// 发送到 TCP 目标
		err = p.SendLog(input)
	case common.MessageGetUpgradeInformation:
		// 发送到 TCP 目标
		err = p.GetUpgradeInformation(input)
	case common.MessageDownloadUrlFile:
		// 发送到 TCP 目标
		err = p.DownloadUrlFile(input)
	case common.MessageGetFileCategories:
		// 发送到 TCP 目标
		err = p.GetFileCategories(input)
	case common.MessageGetFileLevels:
		// 发送到 TCP 目标
		err = p.GetFileLevels(input)
	default:
		return nil, fmt.Errorf("unknown destination: %s", &input.Destination)
	}

	if err != nil || input.Type == common.MessageReport {
		return nil, err
	}

	// 发送响应消息给请求方
	responseMsg := &common.Message{
		SourceID: p.id,
		TargetID: input.SourceID,
		Type:     common.MessageReport, // 标记为同步响应
		Payload:  "Success",            // 可根据实际需求设置响应内容
	}

	return responseMsg, nil
}

func (p *CommunicationPlugin) ChannelLength() int {
	//TODO implement me
	panic("implement me")
}

func (p *CommunicationPlugin) ChannelCapacity() int {
	//TODO implement me
	panic("implement me")
}

func NewCommunicationPlugin(id string, capacity int, httpUrl string, tcpAddress string) *CommunicationPlugin {
	communication := &CommunicationPlugin{
		id:          id,
		messageCh:   make(chan *common.Message, capacity),
		httpHandler: NewHTTPHandler(httpUrl), // 外部服务器 URL
		tcpHandler:  NewTCPHandler(tcpAddress),
	}

	// 初始化 HTTP 服务器
	communication.httpServer = NewHTTPServer(":8701", communication) // 监听地址和端口

	return communication
}

func (p *CommunicationPlugin) PluginID() string {
	return p.id
}

func (p *CommunicationPlugin) MessageChannel() chan *common.Message {
	return p.messageCh
}

func (p *CommunicationPlugin) ParseMsgSyn(Payload interface{}) *interface{} {
	return nil
}

func (p *CommunicationPlugin) HandleAsyncMessage(input chan *common.Message) {
	p.commonCh = input
}

// 启动插件
func (p *CommunicationPlugin) Start() {
	// 启动 HTTP 服务器
	if p.httpServer != nil {
		if err := p.httpServer.Start(); err != nil {
			common.Logger.Errorf("Failed to start HTTP server: %v", err)
			return
		}
	}
	go func() {
		for msg := range p.messageCh {
			switch msg.Type {
			case common.MessageInfo:
				// 处理同步消息，需要发送响应
				if err := p.handleAsyncMessage(msg); err != nil {
					common.Logger.Errorf("Failed to handle sync message: %v", err)
				}
			case common.MessageReport:
				// 处理响应事件
			default:
				// 未知的消息类型，记录警告日志
				common.Logger.Warnf("Unknown message type: %s", msg.Type)
			}
		}
	}()

	// 启动 TCP 服务器，用于接收外部传入的消息
	if p.tcpHandler != nil {
		go p.tcpHandler.Start(p.messageCh)
	}
	//开启管道

	if runtime.GOOS == "windows" {
		p.pipServer = NewPipeServer(`\\.\pipe\docmappipe`)
	} else {
		p.pipServer = NewPipeServer("/tmp/uds_socket")
	}

	if err := p.pipServer.Start(); err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}

// 处理异步消息的方法
func (p *CommunicationPlugin) handleAsyncMessage(msg *common.Message) error {
	var err error
	switch msg.Destination {
	case common.MessageLogin:
		// 发送到 TCP 目标
		err = p.Login(msg)
	case common.MessageAnonymousLogin:
		// 发送到 TCP 目标
		_, err = p.AnonymousLogin(msg)
		if err == nil {
			err = p.tcpHandler.RegisterCallDidline()
		}
	case common.MessageLogout:
		// 发送到 TCP 目标
		err = p.Logout(msg)
	case common.MessageStatOuterAudit:
		// 发送到 TCP 目标
		err = p.StatOuterAudit(msg)
	case common.MessageOuterAudit:
		// 发送到 TCP 目标
		err = p.OuterAudit(msg)
	case common.MessageGetConfig:
		// 发送到 TCP 目标
		err = p.GetConfig(msg)
	case common.MessageSendLog:
		// 发送到 TCP 目标
		err = p.SendLog(msg)
	case common.MessageGetUpgradeInformation:
		// 发送到 TCP 目标
		err = p.GetUpgradeInformation(msg)
	case common.MessageDownloadUrlFile:
		// 发送到 TCP 目标
		err = p.DownloadUrlFile(msg)
	case common.MessageGetFileCategories:
		// 发送到 TCP 目标
		err = p.GetFileCategories(msg)
	case common.MessageGetFileLevels:
		// 发送到 TCP 目标
		err = p.GetFileLevels(msg)
	default:
		common.Logger.Warnf("Unknown destination for sync message: %v", msg.Destination)
		return fmt.Errorf("unknown destination: %v", msg.Destination)
	}

	return err
}
