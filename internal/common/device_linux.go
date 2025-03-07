//go:build linux
// +build linux

package common

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// GetDeviceNames 返回 Linux 上的所有块设备名列表
// 示例：在 Linux 上从 /sys/block 中获取块设备名，如 sda, sdb 等
func GetDeviceNames() []string {
	devNames := []string{}

	files, err := ioutil.ReadDir("/sys/block")
	if err != nil {
		common.Logger.Printf("Failed to read /sys/block: %v", err)
		return devNames
	}

	for _, f := range files {
		devNames = append(devNames, f.Name())
	}
	return devNames
}

// getDeviceID 获取 Linux 设备的唯一标识符
func GetDeviceID() (string, error) {
	// 尝试读取 /etc/machine-id
	machineID, err := readFileTrimmed("/etc/machine-id")
	if err == nil && machineID != "" {
		return machineID, nil
	}

	// 尝试读取 /var/lib/dbus/machine-id
	machineID, err = readFileTrimmed("/var/lib/dbus/machine-id")
	if err == nil && machineID != "" {
		return machineID, nil
	}

	// 如果上述方法失败，尝试通过 DMI 获取主板序列号
	serial, err := getLinuxMotherboardSerial()
	if err == nil && serial != "" {
		return serial, nil
	}

	return "", errors.New("无法获取 Linux 设备ID")
}

// 读取文件并去除空白字符
func readFileTrimmed(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// 通过 DMI 获取 Linux 主板序列号
func getLinuxMotherboardSerial() (string, error) {
	cmd := exec.Command("dmidecode", "-t", "baseboard")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("执行 dmidecode 失败: %v, %s", err, stderr.String())
	}

	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "Serial Number:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				serial := strings.TrimSpace(parts[1])
				if serial != "" && serial != "To be filled by O.E.M." {
					return serial, nil
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("扫描 dmidecode 输出失败: %v", err)
	}

	return "", errors.New("未找到主板序列号")
}
