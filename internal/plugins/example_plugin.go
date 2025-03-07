// Package plugins internal/plugins/example_plugin.go
package plugins

import (
	"docmap-client-proxy-go/internal/common"
	"fmt"
)

type ExamplePlugin struct {
	id        string
	messageCh chan *common.Message
}

func (p *ExamplePlugin) ID() string {
	return p.id
}

func (p *ExamplePlugin) Stop() error {
	//TODO implement me
	panic("implement me")
}

func (p *ExamplePlugin) HandleSyncMessage(input *common.Message) (*common.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ExamplePlugin) HandleAsyncMessage(input chan *common.Message) {
	//TODO implement me
	panic("implement me")
}

func NewExamplePlugin(id string, capacity int) *ExamplePlugin {
	return &ExamplePlugin{
		id:        id,
		messageCh: make(chan *common.Message, capacity), // 根据需要设置合适的容量
	}
}

func (p *ExamplePlugin) PluginID() string {
	return p.id
}

func (p *ExamplePlugin) MessageChannel() chan *common.Message {
	return p.messageCh
}

func (p *ExamplePlugin) ParseMsgSyn(Payload interface{}) *interface{} {
	return nil
}

func (p *ExamplePlugin) Start() {
	go func() {
		for msg := range p.messageCh {
			logger.Info("Plugin %s received message: %v\n", p.id, msg.Payload)
			// 处理消息响应
			if msg.Type == common.MessageInfo && msg.ResponseCh != nil {
				response := &common.Message{
					SourceID: p.id,
					TargetID: msg.SourceID,
					Type:     common.MessageReport,
					Payload:  fmt.Sprintf("Acknowledged message from %s", p.id),
				}
				msg.ResponseCh <- response
			}
		}
	}()
}

func (p *ExamplePlugin) HandleMessage(msg *common.Message) {
	logger.Info("Plugin %s received message from %s", p.PluginID(), msg.SourceID)

	switch payload := msg.Payload.(type) {
	case string:
		logger.Info("Plugin %s received message: %s\n", p.PluginID(), payload)
	// 可以根据需要处理不同类型的消息
	default:
		logger.Warn("Plugin %s received unknown message type", p.PluginID())
	}

	// 如果是同步消息，需要发送响应
	if msg.Type == common.MessageInfo && msg.ResponseCh != nil {
		response := &common.Message{
			SourceID: p.PluginID(),
			TargetID: msg.SourceID,
			Type:     common.MessageReport,
			Payload:  fmt.Sprintf("Acknowledged message: %v", msg.Payload), // 使用 msg.Payload
		}
		msg.ResponseCh <- response
	}
}
