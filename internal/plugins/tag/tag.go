package tag

import (
	"bytes"
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type TagPlugin struct {
	id            string
	messageCh     chan *common.Message
	commonCh      chan *common.Message
	HTTPClient    *http.Client
	ServerAddress string
	iconPathLeft  string
	config        *config.Config
}

func (p *TagPlugin) ID() string {
	return p.id
}

func (p *TagPlugin) Stop() error {
	//TODO implement me
	panic("implement me")
}

func (p *TagPlugin) HandleSyncMessage(msg *common.Message) (*common.Message, error) {
	if msg.Type == common.MessageReport {
		return nil, nil
	}

	report := &common.Message{
		SourceID: p.id,
		TargetID: msg.SourceID,
		Type:     common.MessageReport,
		Payload:  nil,
	}
	var err error
	switch msg.Destination {
	case common.MessageGetFileID:
		report.Payload, err = p.GetFileID(msg)
	case common.MessageGetFileLevel:
		filePath := msg.Payload.(*string)
		report.Payload, err = p.GetFileLevel(*filePath)
	case common.MessageSetFileID:
		report.Payload, err = p.SetFileID(msg)
	case common.MessageSetFileLevel:
		report.Payload, err = p.SetFileLevel(msg)
	case common.MessageDeleteFile:
		filePath := msg.Payload.(*string)
		report.Payload, err = p.DeleteFile(*filePath)
	default:
		common.Logger.Warnf("Unknown destination for sync message: %s", msg.Destination)
		report.Payload = fmt.Errorf("unknown destination: %v", msg.Destination)
		return report, fmt.Errorf("unknown destination: %v", msg.Destination)
	}

	return report, err
}

func (p *TagPlugin) HandleAsyncMessage(input chan *common.Message) {
	//TODO implement me
	p.commonCh = input
}

func (p *TagPlugin) ChannelLength() int {
	//TODO implement me
	panic("implement me")
}

func (p *TagPlugin) ChannelCapacity() int {
	//TODO implement me
	panic("implement me")
}

func NewTagPlugin(id string, capacity int, cfg *config.Config) *TagPlugin {
	return &TagPlugin{
		id:            id,
		messageCh:     make(chan *common.Message, capacity),
		commonCh:      make(chan *common.Message, capacity),
		HTTPClient:    &http.Client{},
		ServerAddress: "http://localhost:8080",
		iconPathLeft:  "/usr/share/icons/docmap-tag/",
		config:        cfg,
	}
}

func (p *TagPlugin) PluginID() string {
	return p.id
}

func (p *TagPlugin) MessageChannel() chan *common.Message {
	return p.messageCh
}

func (p *TagPlugin) ParseMsgSyn(Payload interface{}) *interface{} {
	return nil
}

// Start 启动插件
func (p *TagPlugin) Start() {
	go func() {
		for msg := range p.messageCh {
			switch msg.Type {
			case common.MessageInfo:
				// 处理同步消息，需要发送响应
				if _, err := p.HandleSyncMessage(msg); err != nil {
					common.Logger.Errorf("Failed to handle sync message: %v", err)
				}
			case common.MessageReport:
				// 处理响应事件
			default:
				// 未知的消息类型，记录警告日志
				common.Logger.Warnf("Unknown message type: %s", msg.Type)
			}
		}
	}()
}

// 处理异步消息的方法
func (p *TagPlugin) handleAsyncMessage(msg *common.Message) (*common.Message, error) {
	report := &common.Message{
		SourceID: p.id,
		TargetID: msg.SourceID,
		Type:     common.MessageReport,
		Payload:  nil,
	}
	var err error
	switch msg.Destination {
	case common.MessageGetFileID:
		report.Payload, _ = p.GetFileID(msg)
	case common.MessageGetFileLevel:
		filePath := msg.Payload.(*string)
		report.Payload, err = p.GetFileLevel(*filePath)
	case common.MessageSetFileID:
		report.Payload, err = p.SetFileID(msg)
	case common.MessageSetFileLevel:
		report.Payload, err = p.SetFileLevel(msg)
	case common.MessageDeleteFile:
		filePath := msg.Payload.(*string)
		report.Payload, err = p.DeleteFile(*filePath)
	default:
		common.Logger.Warnf("Unknown destination for sync message: %s", msg.Destination)
		report.Payload = fmt.Errorf("unknown destination: %v", msg.Destination)
		return report, fmt.Errorf("unknown destination: %s", msg.Destination)
	}

	return report, err
}

// GetFileID 处理异步消息的方法
func (p *TagPlugin) GetFileID(msg *common.Message) (*common.FileIdentity, error) {
	filePath := msg.Payload.(*string)
	// 对 fileLevel.FilePath 进行 URL 编码
	encodedFilePath := url.QueryEscape(*filePath)
	getUrl := fmt.Sprintf("%s/identity?filepath=%s", p.ServerAddress, encodedFilePath)
	resp, err := p.HTTPClient.Get(getUrl)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("服务器返回错误: %s", string(body))
	}

	var result struct {
		Family   string `json:"family"`
		Identity string `json:"identity"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}
	FileID := &common.FileIdentity{
		Family:   result.Family,
		Identity: result.Identity,
	}

	return FileID, nil
}

func (p *TagPlugin) GetFileLevel(filePath string) (*common.DocumentClassification, error) {
	// 对 fileLevel.FilePath 进行 URL 编码
	encodedFilePath := url.QueryEscape(filePath)
	getUrl := fmt.Sprintf("%s/get_classification?file_path=%s", p.ServerAddress, encodedFilePath)
	resp, err := p.HTTPClient.Get(getUrl)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("服务器返回错误: %s", string(body))
	}

	var result struct {
		FilePath   string        `json:"file_path"`
		FileSecret common.Secret `json:"secret"`
		Status     string        `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("获取密级属性失败")
	}

	FileID := &common.DocumentClassification{
		FilePath:  result.FilePath,
		FileLevel: result.FileSecret,
	}
	return FileID, nil
}

func (p *TagPlugin) SetFileID(msg *common.Message) (*common.IdentityWriteTask, error) {
	reqData := msg.Payload.(*common.SetIdentityRequest)
	var result struct {
		Family   string `json:"family"`
		Identity string `json:"identity"`
	}
	// 将请求数据编码为JSON
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		common.Logger.Errorf("编码JSON失败: %v", err)
	}
	// 发送POST请求
	postUrl := fmt.Sprintf("%s/identity", p.ServerAddress) // 假设你的服务器地址和接口路径
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		common.Logger.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		common.Logger.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("服务器返回错误: %s", string(body))
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}
	resultId := &common.IdentityWriteTask{
		FilePath: reqData.FilePath,
		Identity: common.FileIdentity{
			Family:   result.Family,
			Identity: result.Identity,
			FileHash: "",
		},
	}
	return resultId, nil
}

func (p *TagPlugin) SetFileLevel(msg *common.Message) (*common.DocumentClassification, error) {
	fileLevel := msg.Payload.(*common.DocumentClassification)
	_, err := os.Stat(fileLevel.FilePath)
	if err != nil {
		p.DeleteFile(fileLevel.FilePath)
		return nil, err
	}
	oldLevel, err := p.GetFileLevel(fileLevel.FilePath)
	if err != nil {
		oldLevel = &common.DocumentClassification{
			FilePath: fileLevel.FilePath,
			FileLevel: common.Secret{
				Level: "0",
				Code:  "0",
			},
		}
	}
	oldLevelValue, err := strconv.Atoi(oldLevel.FileLevel.Level)
	if err != nil {
		// 处理转换错误
	}

	fileLevelValue, err := strconv.Atoi(fileLevel.FileLevel.Level)
	if err != nil {
		// 处理转换错误
	}

	if oldLevelValue < fileLevelValue {
		//设置文件密级
		postUrl := fmt.Sprintf("%s/set_classification", p.ServerAddress)
		//保证自己的密级设置优先级，然后用另一条线对其家族进行整体修改
		payload := map[string]interface{}{
			"file_path": fileLevel.FilePath,
			"secret":    fileLevel.FileLevel,
		}
		data, err := json.Marshal(payload)
		if err != nil {
			return oldLevel, fmt.Errorf("JSON序列化失败: %v", err)
		}

		resp, err := p.HTTPClient.Post(postUrl, "application/json", bytes.NewBuffer(data))
		if err != nil {
			return oldLevel, fmt.Errorf("HTTP POST请求失败: %v", err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("设置密级出错")
		}

		go func() {
			//设置图标
			if common.IsExtensionAllowed(fileLevel.FilePath) {
				ext := filepath.Ext(fileLevel.FilePath)
				// 使用 strings.TrimPrefix 去掉扩展名的点
				extWithoutDot := strings.TrimPrefix(ext, ".")
				iconFile := fmt.Sprintf("%s-%s.png", extWithoutDot, fileLevel.FileLevel.Code)
				iconPath := p.iconPathLeft + iconFile
				err := common.SetFileIcon(fileLevel.FilePath, iconPath)
				if err != nil {
					return
				}
			}
		}()
		//接下来异步操作所有相关文档的密级
		go func() {
			// 对 fileLevel.FilePath 进行 URL 编码
			encodedFilePath := url.QueryEscape(fileLevel.FilePath)
			postUrl = fmt.Sprintf("%s/identity?filepath=%s", "http://localhost:8080", encodedFilePath)
			respEx, err := p.HTTPClient.Get(postUrl)
			if err != nil {
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(respEx.Body)

			if respEx.StatusCode != http.StatusOK {
				return
			}

			var resultEx struct {
				Family   string `json:"family"`
				Identity string `json:"identity"`
			}

			if err = json.NewDecoder(respEx.Body).Decode(&resultEx); err != nil {
				return
			}
			//用Famili查找所有相关文件，然后进行设置
			files, err := p.GetFamilyFiles(resultEx.Family)
			if err != nil {
				return
			}
			//设置所有家族ID密级
			for _, file := range files {
				//判断文件是否存在 不存在 直接掠过并且删除

				if file == fileLevel.FilePath {
					continue
				}
				_, err := os.Stat(file)
				if err != nil {
					p.DeleteFile(file)
					continue
				}
				oldLevelFm, err := p.GetFileLevel(file)
				if err != nil {
					oldLevelFm = &common.DocumentClassification{
						FilePath: fileLevel.FilePath,
						FileLevel: common.Secret{
							Level: "0",
							Code:  "0",
						},
					}
				}

				oldLevelValueFm, err := strconv.Atoi(oldLevelFm.FileLevel.Level)
				if err != nil {
					// 处理转换错误
				}

				fileLevelValueFm, err := strconv.Atoi(fileLevel.FileLevel.Level)
				if err != nil {
					// 处理转换错误
				}
				if oldLevelValueFm > fileLevelValueFm {
					return
				}
				//判断
				postUrl = fmt.Sprintf("%s/set_classification", p.ServerAddress)
				payload := map[string]interface{}{
					"file_path": file,
					"secret":    fileLevel.FileLevel,
				}
				data, err := json.Marshal(payload)
				if err != nil {
					return
				}

				resp, err := p.HTTPClient.Post(postUrl, "application/json", bytes.NewBuffer(data))
				if err != nil {
					return
				}
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {

					}
				}(resp.Body)

				if resp.StatusCode != http.StatusOK {
					continue
				}

				go func() {
					//设置图标
					if common.IsExtensionAllowed(file) {
						ext := filepath.Ext(file)
						// 使用 strings.TrimPrefix 去掉扩展名的点
						extWithoutDot := strings.TrimPrefix(ext, ".")
						iconFile := fmt.Sprintf("%s-%s.png", extWithoutDot, fileLevel.FileLevel.Code)
						iconPath := p.iconPathLeft + iconFile
						err := common.SetFileIcon(file, iconPath)
						if err != nil {
							return
						}
					}
				}()

			}
		}()

	}
	return oldLevel, nil
}

// GetFamilyFiles 根据 familyID 获取所有相关文件的路径
func (p *TagPlugin) GetFamilyFiles(familyID string) ([]string, error) {
	// 构造URL，例如 http://localhost:8080/family/xxxxxx
	url := fmt.Sprintf("%s/family/%s", p.ServerAddress, familyID)

	// 发起 GET 请求
	resp, err := p.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 如果 HTTP 状态码不是 200，返回错误
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("服务器返回错误: %s", string(body))
	}

	// 定义一个结构体来接收服务器的 JSON 响应
	var result struct {
		Files []string `json:"files"`
		Error string   `json:"error,omitempty"`
	}

	// 解析响应体
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	// 如果有错误字段，说明服务器端返回了错误信息
	if result.Error != "" {
		return nil, fmt.Errorf("服务器错误: %s", result.Error)
	}

	// 返回获取到的文件列表
	return result.Files, nil
}

func (p *TagPlugin) DeleteFile(filePath string) (*common.IdentityWriteTask, error) {
	// 对 fileLevel.FilePath 进行 URL 编码
	encodedFilePath := url.QueryEscape(filePath)
	url := fmt.Sprintf("%s/delete_file?filepath=%s", p.ServerAddress, encodedFilePath)
	resp, err := p.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("服务器返回错误: %s", string(body))
	}
	return nil, nil
}
