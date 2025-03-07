package messagebus

import (
	"docmap-client-proxy-go/internal/common"
	"fmt"
	"runtime"
	"sync"
)

type MessageBus struct {
	subscribers  map[string]Plugin
	messageQueue chan *common.Message // 全局异步消息队列
	mu           sync.RWMutex         // 读写锁
}

var (
	MessageBusInstance *MessageBus
	once               sync.Once
)

// GetClientInstance 获取全局唯一的Client实例
func GetMessageBusInstance() *MessageBus {
	once.Do(func() {
		queueSize := 100
		MessageBusInstance = &MessageBus{
			subscribers:  make(map[string]Plugin),
			messageQueue: make(chan *common.Message, queueSize), // 初始化异步消息队列
		}
		common.RegisterInterface("Send", &MessageBusInstance)
	})
	return MessageBusInstance
}

// Register 注册订阅者：每个订阅者都有自己的消息通道
func (bus *MessageBus) Register(subscriber Plugin) {
	bus.mu.Lock() // 使用写锁
	defer bus.mu.Unlock()
	bus.subscribers[subscriber.ID()] = subscriber
}

// Unregister 注销订阅者
func (bus *MessageBus) Unregister(pluginID string) {
	bus.mu.Lock() // 使用写锁
	defer bus.mu.Unlock()

	if _, exists := bus.subscribers[pluginID]; exists {
		delete(bus.subscribers, pluginID)
	} else {
		common.Logger.Warnf("Attempted to unregister non-existent plugin: %s", pluginID)
	}
}

// IsRegistered 检查插件是否已注册
func (bus *MessageBus) IsRegistered(pluginID string) bool {
	bus.mu.RLock() // 使用读锁
	defer bus.mu.RUnlock()

	_, exists := bus.subscribers[pluginID]
	return exists
}

// Start 启动消息分发逻辑
func (bus *MessageBus) Start() {
	numWorkers := runtime.NumCPU() // 根据实际情况调整
	for i := 0; i < numWorkers; i++ {
		go func() {
			for msg := range bus.messageQueue {
				bus.mu.RLock()
				plugin, exists := bus.subscribers[msg.TargetID]
				bus.mu.RUnlock()

				if !exists {
					common.Logger.Warnf("No plugin found for TargetID: %s", msg.TargetID)
					continue
				}
				// 同步消息处理
				go func() {
					_, err := plugin.HandleSyncMessage(msg)
					if err != nil {

					}
				}()

			}
		}()
	}
}

func (bus *MessageBus) Send(msg *common.Message) (*common.Message, error) {
	bus.mu.RLock()
	plugin, exists := bus.subscribers[msg.TargetID]
	bus.mu.RUnlock()

	if !exists {
		common.Logger.Warnf("Target plugin %s not found", msg.TargetID)
		return nil, fmt.Errorf("target plugin %s not found", msg.TargetID)
	}

	response, err := plugin.HandleSyncMessage(msg)
	if err != nil {
		common.Logger.Warnf("Error handling sync message for plugin %s: %v", plugin.ID(), err)
	}
	return response, err
}

func (bus *MessageBus) SubmitAsyncTask(msg *common.Message) {
	bus.messageQueue <- msg
}
