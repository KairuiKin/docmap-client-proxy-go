package register

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	"fmt"
	"time"
)

const (
	MessageLogin   common.MessageType = iota // 登入
	MessageReLogin                           // 重新登入
)

type RegisterPlugin struct {
	id        string
	messageCh chan *common.Message
	commonCh  chan *common.Message
}

func (p *RegisterPlugin) ID() string {
	return p.id
}

func (p *RegisterPlugin) Stop() error {
	//TODO implement me
	panic("implement me")
}

func (p *RegisterPlugin) HandleSyncMessage(input *common.Message) (*common.Message, error) {
	if input.Type == common.MessageReport {
		return nil, nil
	}

	var err error
	switch input.Destination {
	case MessageLogin:
		// 发送到 TCP 目标
		err = p.Login()
	case MessageReLogin:
		// 发送到 TCP 目标
		err = p.ReLogin(input)
	default:
		common.Logger.Warnf("Unknown destination for sync message: %s", input.Destination)
		return nil, fmt.Errorf("unknown destination: %v", input.Destination)
	}

	// 发送响应消息给请求方
	responseMsg := &common.Message{
		SourceID: p.id,
		TargetID: input.SourceID,
		Type:     common.MessageReport, // 标记为同步响应
		Payload:  "Success",            // 可根据实际需求设置响应内容
	}

	return responseMsg, err
}

func (p *RegisterPlugin) HandleAsyncMessage(input chan *common.Message) {
	//TODO implement me
	p.commonCh = input
}

func (p *RegisterPlugin) ChannelLength() int {
	//TODO implement me
	panic("implement me")
}

func (p *RegisterPlugin) ChannelCapacity() int {
	//TODO implement me
	panic("implement me")
}

func NewRegisterPlugin(id string, capacity int) *RegisterPlugin {
	return &RegisterPlugin{
		id:        id,
		messageCh: make(chan *common.Message, capacity),
	}
}

func (p *RegisterPlugin) PluginID() string {
	return p.id
}

func (p *RegisterPlugin) MessageChannel() chan *common.Message {
	return p.messageCh
}

func (p *RegisterPlugin) ParseMsgSyn(Payload interface{}) *interface{} {
	return nil
}

// Start 启动插件
func (p *RegisterPlugin) Start() {
	go func() {
		err := p.Login()
		if err != nil {
			return
		}
		for msg := range p.messageCh {
			switch msg.Type {
			case common.MessageInfo:
				// 处理同步消息，需要发送响应
				if _, err := p.HandleSyncMessage(msg); err != nil {
					common.Logger.Errorf("Failed to handle sync message: %v", err)
				}
			case common.MessageReport:
				// 处理响应事件
			default:
				// 未知的消息类型，记录警告日志
				common.Logger.Warnf("Unknown message type: %v", msg.Type)
			}
		}
	}()
}

func (p *RegisterPlugin) Login() error {
	// 发送同步消息
	asyncMsg := &common.Message{
		SourceID:    p.id,
		TargetID:    "tunnel",
		Destination: common.MessageAnonymousLogin,
		Type:        common.MessageInfo,
		Payload:     "",
	}
	// 创建消息总线
	bus := messagebus.GetMessageBusInstance()
	for {
		_, err := bus.Send(asyncMsg)
		if err == nil {
			// 发送成功，退出循环
			return nil
		}

		// 发送失败，记录错误并等待 5 秒后重试
		common.Logger.Printf("登陆失败: %v。5 秒后重试...", err)
		time.Sleep(5 * time.Second)
	}
}

func (p *RegisterPlugin) ReLogin(msg *common.Message) error {
	return nil
}
