//go:build windows
// +build windows

package common

import (
	"errors"
	"fmt"
	"github.com/StackExchange/wmi"
	"os"
	"strings"
)

// GetDeviceNames 在 Windows 上返回设备名称列表。
// 简单示例：仅返回主机名作为唯一名称。
// 如需获取更详细的设备名称（如磁盘名称、网络接口名称），可通过 WMI 查询。
func GetDeviceNames() []string {
	hostname, err := os.Hostname()
	if err != nil {
		Logger.Errorf("Failed to get hostname: %v", err)
		return []string{"unknown-windows-device"}
	}
	return []string{hostname}
}

// Win32_ComputerSystemProduct 定义 WMI 查询结果的结构体
type Win32_ComputerSystemProduct struct {
	UUID string
}

// getDeviceID 获取 Windows 设备的唯一标识符
func GetDeviceID() (string, error) {
	var dst []Win32_ComputerSystemProduct
	query := "SELECT UUID FROM Win32_ComputerSystemProduct"
	err := wmi.Query(query, &dst)
	if err != nil {
		return "", fmt.Errorf("WMI 查询失败: %v", err)
	}

	if len(dst) == 0 {
		return "", errors.New("未获取到任何 WMI 数据")
	}

	uuid := strings.TrimSpace(dst[0].UUID)
	if uuid == "" || strings.Contains(uuid, "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF") {
		return "", errors.New("获取到的 UUID 无效")
	}
	return uuid, nil
}
