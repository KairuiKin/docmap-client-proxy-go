// cmd/terminal_service/main.go
package main

import (
	"docmap-client-proxy-go/internal/config"
	"docmap-client-proxy-go/internal/messagebus"
	"docmap-client-proxy-go/internal/plugins/event"
	"docmap-client-proxy-go/internal/plugins/register"
	"docmap-client-proxy-go/internal/plugins/tag"
	"docmap-client-proxy-go/internal/plugins/tunnel"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Config 定义配置文件的结构
type Config struct {
	Tunnel struct {
		ServerURL string `yaml:"server_url"`
		Address   string `yaml:"address"`
	} `yaml:"tunnel"`
}

// trimQuotes 去除字符串前后的引号（英文和中文）
func trimQuotes(s string) string {
	// 去除英文引号
	s = strings.TrimPrefix(s, `"`)
	s = strings.TrimSuffix(s, `"`)
	// 去除中文引号
	s = strings.TrimPrefix(s, `“`)
	s = strings.TrimSuffix(s, `”`)
	return s
}

func main() {
	//logger.InitLogger("app.log", logger.INFO)
	//defer logger.CloseLogger()

	// 定义 -f 参数
	configPath := flag.String("f", "config.yaml", "Path to the configuration file")
	flag.Parse()

	// 去除路径前后的引号
	cleanConfigPath := trimQuotes(*configPath)

	// 加载配置
	cfg, err := config.LoadConfig(cleanConfigPath)
	if err != nil {
		fmt.Printf("Error loading config from %s: %v\n", cleanConfigPath, err)
		os.Exit(1)
	}

	// 创建消息总线
	bus := messagebus.GetMessageBusInstance()

	// 使用配置文件中的参数创建插件
	tunnelPlugin := tunnel.NewCommunicationPlugin(
		"tunnel",
		100,
		cfg.Tunnel.ServerURL,
		cfg.Tunnel.Address,
	)
	registerPlugin := register.NewRegisterPlugin("register", 100)
	eventPlugin := event.NewEventPlugin("event", 100, cfg)
	tagPlugin := tag.NewTagPlugin("tag", 100, cfg)

	// 注册插件到消息总线
	bus.Register(tunnelPlugin)
	bus.Register(registerPlugin)
	bus.Register(eventPlugin)
	bus.Register(tagPlugin)
	bus.Start()
	// 启动插件
	tunnelPlugin.Start()
	registerPlugin.Start()
	eventPlugin.Start()
	tagPlugin.Start()
	// 设置信号接收器来处理 Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press Ctrl+C to exit...")
	<-stop // 阻塞，直到接收到信号

	// 执行清理操作
	fmt.Println("\nShutting down gracefully...")

	// 遍历所有插件，判断它们是否存活并进行注销
	pluginsToUnregister := []string{"example-plugin-1", "tunnel"}
	for _, pluginID := range pluginsToUnregister {
		if bus.IsRegistered(pluginID) {
			bus.Unregister(pluginID)
			fmt.Printf("Plugin %s unregistered successfully.\n", pluginID)
		} else {
			fmt.Printf("Plugin %s is already unregistered.\n", pluginID)
		}
	}

	fmt.Println("All active plugins unregistered. Exiting.")

}
