package messagebus

import "docmap-client-proxy-go/internal/common"

const (
	MessageInfo   common.MessageType = iota // 同步消息
	MessageReport                           //响应
)

type Subscriber interface {
	PluginID() string
	MessageChannel() chan *common.Message
	ParseMsgSyn(Payload interface{}) *interface{}
}

type Plugin interface {
	ID() string
	Start()
	Stop() error
	HandleSyncMessage(input *common.Message) (*common.Message, error) // 同步处理消息
	HandleAsyncMessage(input chan *common.Message)                    // 异步处理消息
}
