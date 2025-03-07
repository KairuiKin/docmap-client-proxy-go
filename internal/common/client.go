package common

import (
	"bytes"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

// ClientInfo 表示一个 Application 和一个无符号整数的组合
type ClientInfo struct {
	Application *__.Application
	ID          uint
}

// Client 表示一个 Socket 客户端
type Client struct {
	DeviceID      string
	DeviceName    []string
	AccessToken   string
	UserId        string
	seedId        uint64
	IsStart       bool
	CurrentUser   *user.User
	ClientOS      *__.OS
	CurrentClient *__.Client
	ClientsMap    sync.Map // 使用 sync.Map
}

var (
	clientInstance *Client
	once           sync.Once
)

// GetClientInstance 获取全局唯一的Client实例
func GetClientInstance() *Client {
	once.Do(func() {
		current, err := user.Current()
		if err != nil {
			return
		}
		deivceId, _ := GetDeviceID()
		clientInstance = &Client{
			DeviceID:      deivceId,
			DeviceName:    GetDeviceNames(),
			AccessToken:   "",
			UserId:        "",
			seedId:        0,
			IsStart:       false,
			CurrentUser:   current,
			ClientOS:      &__.OS{},
			CurrentClient: &__.Client{},
		}
		osVersion, _ := getSystemVersion()
		clientInstance.ClientOS.OsName = wrapperspb.String(runtime.GOOS)
		clientInstance.ClientOS.OsArch = wrapperspb.String(runtime.GOARCH)
		clientInstance.ClientOS.OsVersion = wrapperspb.String(osVersion)
		info, s, err := GetADDomainInfo()
		if err != nil {
			return
		}
		clientInstance.ClientOS.InAd = wrapperspb.Bool(info)
		clientInstance.ClientOS.AdName = wrapperspb.String(s)
		clientInstance.CurrentClient.Name = wrapperspb.String("docmap-monitor")
		clientInstance.CurrentClient.Version = wrapperspb.String("0.1.1.68")
		clientInstance.CurrentClient.Build = wrapperspb.UInt64(0)

	})
	return clientInstance
}

func (c *Client) GetMessageId() uint64 {
	return atomic.AddUint64(&c.seedId, 1)
}

// getSystemVersion 获取系统版本
func getSystemVersion() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "ver")
	case "linux":
		cmd = exec.Command("uname", "-r")
	case "darwin":
		cmd = exec.Command("sw_vers", "-productVersion")
	default:
		return "", fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		// 将 output 转换为 UTF-8 编码
		reader := transform.NewReader(bytes.NewReader(output), simplifiedchinese.GBK.NewDecoder())
		decodedOutput, err := io.ReadAll(reader)
		if err != nil {
			return "", err
		}
		output = decodedOutput
	}
	return strings.TrimSpace(string(output)), nil
}

// 添加客户端
func (c *Client) AddClient(key string, whoImI *model.WhoAmI, id uint) {
	var app __.Application
	app.Url = wrapperspb.String("")
	app.DisplayName = wrapperspb.String(filepath.Base(whoImI.App))
	app.ExecutableName = wrapperspb.String(whoImI.App)
	app.ClassificationType = append(app.ClassificationType, wrapperspb.String(""))
	app.SecurityLevel = wrapperspb.String("common")
	app.ExeSha1 = wrapperspb.String("")
	c.ClientsMap.Store(key, ClientInfo{
		Application: &app,
		ID:          id,
	})
}

// 获取客户端
func (c *Client) GetClient(key string) (ClientInfo, bool) {
	value, exists := c.ClientsMap.Load(key)
	if !exists {
		return ClientInfo{}, false
	}
	return value.(ClientInfo), true
}

// 删除客户端
func (c *Client) DeleteClient(key string) {
	c.ClientsMap.Delete(key)
}
