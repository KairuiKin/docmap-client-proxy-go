package tunnel

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

// SecurityLevelRequest 定义了接收的 JSON 请求结构

// SecurityLevelResponse 定义了返回的 JSON 响应结构
type SecurityLevelResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HTTPServer 定义了 HTTP 服务器的结构
type HTTPServer struct {
	server        *http.Server
	communication *CommunicationPlugin
	wg            sync.WaitGroup
}

// NewHTTPServer 创建一个新的 HTTPServer
func NewHTTPServer(address string, communication *CommunicationPlugin) *HTTPServer {
	return &HTTPServer{
		communication: communication,
		server: &http.Server{
			Addr: address,
		},
	}
}

// Start 启动 HTTP 服务器
func (h *HTTPServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/setSecurityLevel", h.handleSetSecurityLevel)

	h.server.Handler = mux

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		common.Logger.Infof("Starting HTTP server at %s", h.server.Addr)
		if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			common.Logger.Errorf("HTTP server ListenAndServe: %v", err)
		}
	}()

	return nil
}

// Stop 优雅地停止 HTTP 服务器
func (h *HTTPServer) Stop() error {
	common.Logger.Infof("Shutting down HTTP server at %s", h.server.Addr)
	if err := h.server.Close(); err != nil {
		return fmt.Errorf("failed to close HTTP server: %v", err)
	}
	h.wg.Wait()
	return nil
}

// handleSetSecurityLevel 处理 /api/setSecurityLevel 路径的 POST 请求
func (h *HTTPServer) handleSetSecurityLevel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req common.SecurityLevelRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		common.Logger.Warnf("Error decoding JSON: %v", err)
		resp := SecurityLevelResponse{
			Status:  "error",
			Message: "Invalid JSON format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	common.Logger.Infof("Received security level %d for file %s", req.SecurityLevel, req.FilePath)

	// 构建消息并发送到 CommunicationPlugin 的 messageCh
	msg := &common.Message{
		SourceID:    "tunnel",
		TargetID:    "event",
		Type:        common.MessageLogin,
		Destination: common.MessageTag,
		Payload: common.SecurityLevelRequest{
			FilePath:      req.FilePath,
			SecurityLevel: req.SecurityLevel,
		},
	}
	bus := messagebus.GetMessageBusInstance()
	bus.SubmitAsyncTask(msg)
}
