// internal/plugins/manager.go
package plugins

import (
	"docmap-client-proxy-go/internal/logger"
	"docmap-client-proxy-go/internal/messagebus"
	"fmt"
	"sync"
)

type Manager struct {
	plugins map[string]messagebus.Plugin
	bus     *messagebus.MessageBus
	mu      sync.Mutex
}

func NewManager(bus *messagebus.MessageBus) *Manager {
	return &Manager{
		plugins: make(map[string]messagebus.Plugin),
		bus:     bus,
	}
}

func (m *Manager) Load(name string, plugin messagebus.Plugin, config interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.plugins[name]; exists {
		return fmt.Errorf("plugin %s already loaded", name)
	}
	plugin.Start()
	m.plugins[name] = plugin
	logger.Info("Plugin %s loaded with ID %s", name, plugin.ID())
	return nil
}

func (m *Manager) Unload(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	plugin, exists := m.plugins[name]
	if !exists {
		return fmt.Errorf("plugin %s not found", name)
	}
	if err := plugin.Stop(); err != nil {
		return err
	}
	delete(m.plugins, name)
	logger.Info("Plugin %s unloaded", name)
	return nil
}

// 获取插件实例
func (m *Manager) GetPluginByID(id string) (messagebus.Plugin, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, plugin := range m.plugins {
		if plugin.ID() == id {
			return plugin, true
		}
	}
	return nil, false
}

// Plugins 方法返回所有已加载的插件
func (m *Manager) Plugins() []messagebus.Plugin {
	m.mu.Lock()
	defer m.mu.Unlock()
	plugins := make([]messagebus.Plugin, 0, len(m.plugins))
	for _, plugin := range m.plugins {
		plugins = append(plugins, plugin)
	}
	return plugins
}
