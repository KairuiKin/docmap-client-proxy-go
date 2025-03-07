package tunnel

import (
	"bytes"
	"docmap-client-proxy-go/internal/common"
	__ "docmap-client-proxy-go/internal/proto"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type HTTPHandler struct {
	URL    string
	client *http.Client
}

func NewHTTPHandler(url string) *HTTPHandler {
	return &HTTPHandler{
		URL:    url,
		client: &http.Client{Timeout: 6 * time.Second},
	}
}

// Login 登入(/dev-auth-api/login/account)
func (p *CommunicationPlugin) Login(msg *common.Message) error {
	common.Logger.Infof("Handling login request from: %s", msg.SourceID)

	return nil
}

// AnonymousLogin 匿名登录（/dev-docmap-agent-api/secure/login）
func (p *CommunicationPlugin) AnonymousLogin(msg *common.Message) (*Response, error) {
	client := common.GetClientInstance()

	// 构建请求体
	loginRequest := struct {
		DeviceID    string   `json:"deviceId"`
		DeviceNames []string `json:"deviceNames"`
	}{
		DeviceID:    client.DeviceID,
		DeviceNames: client.DeviceName,
	}

	// 序列化请求体为 JSON
	jsonData, err := json.Marshal(loginRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal login request: %v", err)
	}
	// 创建 HTTP 请求
	url := p.httpHandler.URL + "/dev-docmap-agent-api/secure/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request for login: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送 HTTP 请求
	response, err := p.httpHandler.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request for login: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login request returned non-OK status: %d", response.StatusCode)
	}

	// 解析响应体
	var webResponse Response

	err = json.NewDecoder(response.Body).Decode(&webResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse login response: %v", err)
	}

	if webResponse.Status != 0 {
		return nil, fmt.Errorf("login failed with status code: %d", webResponse.Status)
	}
	//记录匿名登陆的信息
	Client := common.GetClientInstance()
	Client.AccessToken = webResponse.Result.AccessToken
	Client.UserId = webResponse.Result.UserID
	return &webResponse, nil
}

// Logout 登出(/dev-auth-api/logout)
func (p *CommunicationPlugin) Logout(msg *common.Message) error {
	common.Logger.Infof("Handling logout request for: %s", msg.SourceID)
	// 登出的具体逻辑
	return nil
}

// StatOuterAudit 外发申请(/dev-docmap-agent-api/outerAudit/startOuterAudit)
func (p *CommunicationPlugin) StatOuterAudit(msg *common.Message) error {
	common.Logger.Infof("Handling send log request for: %s", msg.SourceID)
	// 发送日志的具体逻辑
	return nil
}

// OuterAudit 外发申请(/dev-docmap-agent-api/outerAudit/startOuterAudit)
func (p *CommunicationPlugin) OuterAudit(msg *common.Message) error {
	common.Logger.Infof("Handling send log request for: %s", msg.SourceID)
	// 发送日志的具体逻辑
	return nil
}

// GetConfig 获取配置（/dev-docmap-agent-api/get）
func (p *CommunicationPlugin) GetConfig(msg *common.Message) error {
	common.Logger.Infof("Handling send log request for: %s", msg.SourceID)
	// 发送日志的具体逻辑
	return nil
}

// SendLog 发送日志（/dev-docmap-agent-api/log/send）
func (p *CommunicationPlugin) SendLog(msg *common.Message) error {
	var eventLog __.EventLog
	err := proto.Unmarshal(msg.Payload.([]byte), &eventLog)
	if err != nil {
		return err
	}
	var jsonEvent []byte
	switch eventLog.EventName.GetValue() {
	case common.EventInfoEventNameNewFile:
		eventCreate := eventLog.Event.Create
		jsonEvent, err = protojson.Marshal(eventCreate)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameCopy:
		eventCopy := eventLog.Event.Copy
		jsonEvent, err = protojson.Marshal(eventCopy)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameSaveAs:
		eventSaveAs := eventLog.Event.SaveAs
		jsonEvent, err = protojson.Marshal(eventSaveAs)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameZip:
		eventZip := eventLog.Event.Zip
		jsonEvent, err = protojson.Marshal(eventZip)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameUnzip:
		eventUnzip := eventLog.Event.Unzip
		jsonEvent, err = protojson.Marshal(eventUnzip)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameUpload:
	case common.EventInfoEventNameDownload:
	case common.EventInfoEventNameEdit:
		eventEdit := eventLog.Event.Edit
		jsonEvent, err = protojson.Marshal(eventEdit)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameMove:
		eventMove := eventLog.Event.Move
		jsonEvent, err = protojson.Marshal(eventMove)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameRename:
		eventRename := eventLog.Event.Rename
		jsonEvent, err = protojson.Marshal(eventRename)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameDelete:
		eventDelete := eventLog.Event.Delete
		jsonEvent, err = protojson.Marshal(eventDelete)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameRecycle:
		eventRecycle := eventLog.Event.Recycle
		jsonEvent, err = protojson.Marshal(eventRecycle)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameRestore:
		eventRestore := eventLog.Event.Restore
		jsonEvent, err = protojson.Marshal(eventRestore)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameChangeAttr:
	case common.EventInfoEventNameWriteTag:
		eventWriteTag := eventLog.Event.WriteTag
		jsonEvent, err = protojson.Marshal(eventWriteTag)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	case common.EventInfoEventNameRenameFolder:
		eventRenameFolder := eventLog.Event.RenameFolder
		jsonEvent, err = protojson.Marshal(eventRenameFolder)
		if err != nil {
			common.Logger.Errorf("%v", err)
			return err
		}
	default:
		common.Logger.Infof("Unknown event name: %s\n", eventLog.EventName.GetValue())
	}
	eventLog.Event = nil

	client := common.GetClientInstance()
	// 构造请求数据
	requestData := map[string]interface{}{
		"deviceId":    client.DeviceID,
		"deviceNames": client.DeviceName,
	}
	root := map[string]interface{}{}
	root["device"] = requestData
	root["logs"] = []interface{}{}

	jsonData, err := protojson.Marshal(&eventLog)
	if err != nil {
		common.Logger.Errorf("无法将 Protobuf 转为 JSON: %v", err)
		return err
	}

	// 定义一个 map 以存储解析后的 JSON 数据
	var resultEventMap map[string]interface{}

	err = json.Unmarshal(jsonEvent, &resultEventMap)
	if err != nil {
		common.Logger.Errorf("无法将 JSON 转为 map: %v", err)
		return err
	}

	var resultMap map[string]interface{}
	// 将 JSON 字符串解码为 map
	err = json.Unmarshal(jsonData, &resultMap)
	if err != nil {
		common.Logger.Errorf("无法将 JSON 转为 map: %v", err)
		return err
	}
	resultMap["event"] = resultEventMap
	root["logs"] = append(root["logs"].([]interface{}), resultMap)

	// 将请求数据编码为 JSON
	jsonData, err = json.Marshal(root)
	if err != nil {
		return err
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", p.httpHandler.URL+"/dev-docmap-agent-api/log/send", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 添加自定义请求头，例如 Authorization、User-Agent 等
	req.Header.Set("ui", client.UserId)                                          // 设置授权 token
	req.Header.Set("at", client.AccessToken)                                     // 设置 User-Agent
	req.Header.Set("tm", strconv.FormatUint(uint64(eventLog.LogTime.Value), 10)) // 其他自定义的请求头
	// 发送请求
	resp, err := p.httpHandler.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		common.Logger.Errorf("数据发送失败，响应状态码: %d,响应内容: %s", resp.StatusCode, string(body))
		return fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return nil
}

// GetUpgradeInformation 获取升级信息(/dev-docmap-agent-api/upgrade)
func (p *CommunicationPlugin) GetUpgradeInformation(msg *common.Message) error {
	common.Logger.Infof("Handling send log request for: %s", msg.SourceID)
	// 发送日志的具体逻辑
	return nil
}

// DownloadUrlFile 下载指定的文件
func (p *CommunicationPlugin) DownloadUrlFile(msg *common.Message) error {
	common.Logger.Infof("Handling send log request for: %s", msg.SourceID)
	// 发送日志的具体逻辑
	return nil
}

// GetFileCategories 获取文件类别范围
func (p *CommunicationPlugin) GetFileCategories(msg *common.Message) error {
	common.Logger.Infof("Handling send log request for: %s", msg.SourceID)
	// 发送日志的具体逻辑
	return nil
}

// GetFileLevels 获取文件密级范围
func (p *CommunicationPlugin) GetFileLevels(msg *common.Message) error {
	common.Logger.Infof("Handling send log request for: %s", msg.SourceID)
	// 发送日志的具体逻辑
	return nil
}
