package common

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetADDomainInfo 检查是否加入AD域并获取域名
func GetADDomainInfo() (bool, string, error) {
	// 使用 wmic 命令获取域信息
	cmd := exec.Command("wmic", "computersystem", "get", "domain")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, "", err
	}

	// 解析输出
	output := strings.TrimSpace(out.String())
	lines := strings.Split(output, "\n")
	if len(lines) < 2 {
		return false, "", errors.New("无法解析wmic命令的输出")
	}

	domain := strings.TrimSpace(lines[1])
	if strings.EqualFold(domain, "WORKGROUP") || strings.EqualFold(domain, "DOMAIN") {
		// 系统未加入域
		return false, "", nil
	}

	if domain != "" {
		return true, domain, nil
	}

	return false, "", nil
}

// setFileIcon 使用 gio 命令设置文件的自定义图标
func SetFileIcon(filePath string, iconPath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to touch file: %w", err)
	}
	// Step 1: 使用 gio 命令为文件设置自定义图标
	cmd := exec.Command("gio", "set", filePath, "metadata::custom-icon", iconPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set icon: %w", err)
	}

	// Step 2: 执行 touch 命令更新文件时间戳
	cmdTouch := exec.Command("touch", filePath)
	if err := cmdTouch.Run(); err != nil {
		return fmt.Errorf("failed to touch file: %w", err)
	}

	return nil
}
