// internal/registry/registry.go
package registry

import (
	"docmap-client-proxy-go/internal/logger"
	"fmt"
	"sync"
)

type User struct {
	Username string
	Password string
}

type Service struct {
	users map[string]*User
	mu    sync.RWMutex
}

func NewService() *Service {
	return &Service{
		users: make(map[string]*User),
	}
}

func (s *Service) Register(username, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[username]; exists {
		return fmt.Errorf("user %s already exists", username)
	}
	s.users[username] = &User{Username: username, Password: password}
	logger.Info("User registered: %s", username)
	return nil
}

func (s *Service) Unregister(username string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[username]; !exists {
		return fmt.Errorf("user %s not found", username)
	}
	delete(s.users, username)
	logger.Info("User unregistered: %s", username)
	return nil
}
