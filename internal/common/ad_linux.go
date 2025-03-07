package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetADDomainInfo 在非Windows系统下不执行任何操作
func GetADDomainInfo() (bool, string, error) {
	return false, "", nil
}

type RequestData struct {
	FilePath string `json:"filePath"`
	IconPath string `json:"iconPath"`
}

// setFileIcon 使用 gio 命令设置文件的自定义图标
func SetFileIcon(filePath string, iconPath string) error {
	// 请求的数据
	data := RequestData{
		FilePath: filePath,
		IconPath: iconPath,
	}

	// 将结构体编码成 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		common.Logger.Errorf("Error marshaling data: %v", err)
	}

	// 创建 POST 请求
	url := "http://localhost:8081/set-icon"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		common.Logger.Errorf("Error creating request: %v", err)
	}

	// 设置请求头 Content-Type 为 application/json
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		common.Logger.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Response Status: %v ", resp.Status)
	}

	return nil
}
