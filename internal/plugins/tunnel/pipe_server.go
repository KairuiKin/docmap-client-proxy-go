package tunnel

import (
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
)

// PipeServer 是跨平台的管道服务器接口
type PipeServer interface {
	Start() error
	Stop() error
	Send(clientID string, message []byte) error
}

// 基础实现：线程安全管理
type PipeServerBase struct {
	mu      sync.Mutex
	running bool
}

// Start 启动服务
func (ps *PipeServerBase) Start() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.running {
		return errors.New("server already running")
	}
	ps.running = true
	return nil
}

// Stop 停止服务
func (ps *PipeServerBase) Stop() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.running = false
	return nil
}

func DecodeHeader(headerData []byte) (uint32, error) {
	if len(headerData) != 4 {
		return 0, fmt.Errorf("invalid header length: expected 4, got %d", len(headerData))
	}
	return binary.BigEndian.Uint32(headerData), nil
}
