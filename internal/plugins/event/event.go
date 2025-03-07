package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/config"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"runtime"
	"time"
)

type EventPlugin struct {
	id        string
	messageCh chan *common.Message
	commonCh  chan *common.Message
	config    *config.Config
}

func (p *EventPlugin) ID() string {
	return p.id
}

func (p *EventPlugin) Stop() error {
	//TODO implement me
	panic("implement me")
}

func (p *EventPlugin) HandleSyncMessage(input *common.Message) (*common.Message, error) {
	if input.Type == common.MessageReport {
		return nil, nil
	}
	var err error
	switch input.Destination {
	case common.MessageLog:
		var parsy model.ExchangeMessage
		err := proto.Unmarshal(input.Payload.([]byte), &parsy)
		if err != nil {
			common.Logger.Errorf("Received message faild %v", err)
			return nil, err
		}
		var event model.ReportActionEvent
		err = anypb.UnmarshalTo(parsy.Payload, &event, proto.UnmarshalOptions{})
		if err != nil {
			return nil, err
		}
		err = p.ParseActionEvent(parsy.Me.ClientInstanceId, &event)
		if err != nil {
			err = p.ParseActionEventEx(parsy.Me.ClientInstanceId, &event)
			if err != nil {
				return nil, err
			}
		}

	case common.MessageTag:
		req, ok := input.Payload.(common.SecurityLevelRequest)
		if !ok {
			common.Logger.Errorf("Payload is not SecurityLevelRequest, got %T", input.Payload)
			return nil, fmt.Errorf("invalid payload for tag message")
		}
		p.ParseTagEvent("", &req)
		common.Logger.Infof("Tag message: FilePath=%s, SecurityLevel=%v", req.FilePath, req.SecurityLevel)
	default:
		common.Logger.Warnf("Unknown destination for sync message: %s", input.Destination)
		return nil, fmt.Errorf("unknown destination: %v", input.Destination)
	}

	// 发送响应消息给请求方
	responseMsg := &common.Message{
		SourceID: p.ID(),
		TargetID: input.SourceID,
		Type:     common.MessageReport, // 标记为同步响应
		Payload:  "Success",            // 可根据实际需求设置响应内容
	}

	return responseMsg, err
}

func (p *EventPlugin) HandleAsyncMessage(input chan *common.Message) {
	//TODO implement me
	p.commonCh = input
}

func (p *EventPlugin) ChannelLength() int {
	//TODO implement me
	panic("implement me")
}

func (p *EventPlugin) ChannelCapacity() int {
	//TODO implement me
	panic("implement me")
}

func NewEventPlugin(id string, capacity int, cfg *config.Config) *EventPlugin {
	return &EventPlugin{
		id:        id,
		messageCh: make(chan *common.Message, capacity),
		config:    cfg,
	}
}

func (p *EventPlugin) PluginID() string {
	return p.id
}

func (p *EventPlugin) MessageChannel() chan *common.Message {
	return p.messageCh
}

func (p *EventPlugin) ParseMsgSyn(Payload interface{}) *interface{} {
	return nil
}

// Start 启动插件
func (p *EventPlugin) Start() {
	numWorkers := runtime.NumCPU() // 根据实际情况调整
	for i := 0; i < numWorkers; i++ {
		go p.worker()
	}
}

func (p *EventPlugin) worker() {
	for msg := range p.messageCh {
		switch msg.Type {
		case common.MessageInfo:
			if _, err := p.HandleSyncMessage(msg); err != nil {
				common.Logger.Errorf("处理同步消息失败: %v", err)
			}
		case common.MessageReport:
			// 处理响应事件
		default:
			common.Logger.Warnf("未知的消息类型: %v", msg.Type)
		}
	}
}

// Start 启动插件
func (p *EventPlugin) ParseActionEvent(clientID string, event *model.ReportActionEvent) error {
	client := common.GetClientInstance()
	clientInfo, _ := client.GetClient(clientID)
	millis := int64(event.GetEventTime())
	var eventDetail proto.Message

	logBase := p.FillLog(clientInfo.Application, event.EventName, "allow", millis, millis)
	switch event.EventName {
	case common.EventInfoEventNameNewFile:
		var action model.EventCreate
		eventDetail = &action
	case common.EventInfoEventNameCopy:
		var action model.EventCopy
		eventDetail = &action
	case common.EventInfoEventNameSaveAs:
		var action model.EventSaveAs
		eventDetail = &action
	case common.EventInfoEventNameZip:
		var action model.EventZip
		eventDetail = &action
	case common.EventInfoEventNameUnzip:
		var action model.EventUnzip
		eventDetail = &action
	case common.EventInfoEventNameUpload:
	case common.EventInfoEventNameDownload:
	case common.EventInfoEventNameEdit:
		var action model.EventEdit
		eventDetail = &action
	case common.EventInfoEventNameMove:
		var action model.EventMove
		eventDetail = &action
	case common.EventInfoEventNameRename:
		var action model.EventRename
		eventDetail = &action
	case common.EventInfoEventNameDelete:
		var action model.EventDelete
		eventDetail = &action
	case common.EventInfoEventNameRecycle:
		var action model.EventRecycle
		eventDetail = &action
	case common.EventInfoEventNameRestore:
		var action model.EventRestore
		eventDetail = &action
	case common.EventInfoEventNameChangeAttr:
	case common.EventInfoEventNameRenameFolder:
		var action model.EventRenameFolder
		eventDetail = &action
	default:
		return fmt.Errorf("unknown event: %s", event.EventName)
	}

	if err := anypb.UnmarshalTo(event.EventDetail, eventDetail, proto.UnmarshalOptions{}); err != nil {
		return err
	}

	common.Logger.Debugf("文件事件%s,内容%s", event.EventName, event.EventDetail.String())
	// 填充具体的事件详情
	p.populateEventDetail(logBase, event.EventName, eventDetail)

	// 序列化日志
	marshaledLog, err := proto.Marshal(logBase)
	if err != nil {
		return err
	}

	// 创建并提交异步消息
	asyncMsg := &common.Message{
		SourceID:    p.id,
		TargetID:    "tunnel",
		Destination: common.MessageSendLog,
		Type:        common.MessageInfo,
		Payload:     marshaledLog,
	}
	messagebus.GetMessageBusInstance().SubmitAsyncTask(asyncMsg)
	return nil
}

// Start 启动插件
func (p *EventPlugin) ParseActionEventEx(clientID string, event *model.ReportActionEvent) error {
	client := common.GetClientInstance()
	clientInfo, _ := client.GetClient(clientID)
	millis := int64(event.GetEventTime())
	var eventDetail proto.Message
	logBase := p.FillLog(clientInfo.Application, event.EventName, "allow", millis, millis)
	switch event.EventName {
	case common.EventInfoEventNameMoveFolder:
		var action model.EventMoveFolder
		eventDetail = &action
	case common.EventInfoEventNameDeleteFolder:
		var action model.EventDeleteFolder
		eventDetail = &action
	case common.EventInfoEventNameRecycleFolder:
		var action model.EventRecycleFolder
		eventDetail = &action
	case common.EventInfoEventNameRestoreFolder:
		var action model.EventRestoreFolder
		eventDetail = &action
	default:
		return fmt.Errorf("unknown event: %s", event.EventName)
	}

	if err := anypb.UnmarshalTo(event.EventDetail, eventDetail, proto.UnmarshalOptions{}); err != nil {
		return err
	}
	var logAll []*__.EventLog
	// 填充具体的事件详情
	logAll = p.populateEventDetailEx(logBase, event.EventName, eventDetail)
	for _, log := range logAll {
		logId, _ := common.DMRandomUuid()
		log.LogId = wrapperspb.String(logId)
		// 序列化日志
		marshaledLog, err := proto.Marshal(log)
		if err != nil {
			return err
		}

		// 创建并提交异步消息
		asyncMsg := &common.Message{
			SourceID:    p.id,
			TargetID:    "tunnel",
			Destination: common.MessageSendLog,
			Type:        common.MessageInfo,
			Payload:     marshaledLog,
		}
		messagebus.GetMessageBusInstance().SubmitAsyncTask(asyncMsg)
	}
	return nil
}

func (p *EventPlugin) populateEventDetail(logBase *__.EventLog, eventName string, detail proto.Message) {
	switch eventName {
	case common.EventInfoEventNameNewFile:
		action, ok := detail.(*model.EventCreate)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventCreate，实际 %T", detail)
			return
		}
		createFile, _ := p.eventCreate(action)
		logBase.Event = &__.Event{Create: createFile}
	case common.EventInfoEventNameCopy:
		action, ok := detail.(*model.EventCopy)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventCopy，实际 %T", detail)
			return
		}
		copyFile, _ := p.eventCopy(action)
		logBase.Event = &__.Event{Copy: copyFile}
	case common.EventInfoEventNameSaveAs:
		action, ok := detail.(*model.EventSaveAs)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventSaveAs，实际 %T", detail)
			return
		}
		saveAs, _ := p.eventSaveAs(action)
		logBase.Event = &__.Event{SaveAs: saveAs}
	case common.EventInfoEventNameZip:
		action, ok := detail.(*model.EventZip)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventZip，实际 %T", detail)
			return
		}
		zip, _ := p.eventZip(action)
		logBase.Event = &__.Event{Zip: zip}
	case common.EventInfoEventNameUnzip:
		action, ok := detail.(*model.EventUnzip)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventUnzip，实际 %T", detail)
			return
		}
		unzip, _ := p.eventUnzip(action)
		logBase.Event = &__.Event{Unzip: unzip}
	case common.EventInfoEventNameUpload:
		// TODO: 处理 Upload 事件
	case common.EventInfoEventNameDownload:
		// TODO: 处理 Download 事件
	case common.EventInfoEventNameEdit:
		action, ok := detail.(*model.EventEdit)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventEdit，实际 %T", detail)
			return
		}
		edit, _ := p.eventEdit(action)
		logBase.Event = &__.Event{Edit: edit}
	case common.EventInfoEventNameMove:
		action, ok := detail.(*model.EventMove)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventMove，实际 %T", detail)
			return
		}
		moveFile, _ := p.eventMove(action)
		logBase.Event = &__.Event{Move: moveFile}
	case common.EventInfoEventNameRename:
		action, ok := detail.(*model.EventRename)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventRename，实际 %T", detail)
			return
		}
		renameFile, _ := p.eventRename(action)
		logBase.Event = &__.Event{Rename: renameFile}
	case common.EventInfoEventNameDelete:
		action, ok := detail.(*model.EventDelete)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventDelete，实际 %T", detail)
			return
		}
		deleteFile, _ := p.eventDelete(action)
		logBase.Event = &__.Event{Delete: deleteFile}
	case common.EventInfoEventNameRecycle:
		recycleFile, _ := p.eventRecycle(detail.(*model.EventRecycle))
		logBase.Event = &__.Event{Recycle: recycleFile}
	case common.EventInfoEventNameRestore:
		restoreFile, _ := p.eventRestore(detail.(*model.EventRestore))
		logBase.Event = &__.Event{Restore: restoreFile}
	case common.EventInfoEventNameChangeAttr:
	case common.EventInfoEventNameRenameFolder:
		renameFolder, _ := p.eventRenameFolder(detail.(*model.EventRenameFolder))
		logBase.Event = &__.Event{RenameFolder: renameFolder}
	default:
		return
	}
}

func (p *EventPlugin) populateEventDetailEx(logBase *__.EventLog, eventName string, detail proto.Message) []*__.EventLog {
	var logAll []*__.EventLog
	switch eventName {
	case common.EventInfoEventNameMoveFolder:
		action, ok := detail.(*model.EventMoveFolder)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventCreate，实际 %T", detail)
			return logAll
		}
		logBase.EventName = wrapperspb.String(common.EventInfoEventNameMove)
		eventMoves, _ := p.eventMoveFolder(action)
		for _, move := range eventMoves {
			// 克隆 logBase
			logCopy := proto.Clone(logBase).(*__.EventLog)
			logCopy.Event = &__.Event{Move: move}
			logAll = append(logAll, logCopy)
		}
	case common.EventInfoEventNameDeleteFolder:
	case common.EventInfoEventNameRecycleFolder:
		action, ok := detail.(*model.EventRecycleFolder)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventCreate，实际 %T", detail)
			return logAll
		}
		logBase.EventName = wrapperspb.String(common.EventInfoEventNameRecycle)
		eventRecycles, _ := p.eventRecycleFolder(action)
		for _, recycle := range eventRecycles {
			logCopy := proto.Clone(logBase).(*__.EventLog)
			logCopy.Event = &__.Event{Recycle: recycle}
			logAll = append(logAll, logCopy)
		}
	case common.EventInfoEventNameRestoreFolder:
		action, ok := detail.(*model.EventRestoreFolder)
		if !ok {
			common.Logger.Errorf("类型断言失败，预期 *model.EventCreate，实际 %T", detail)
			return logAll
		}
		logBase.EventName = wrapperspb.String(common.EventInfoEventNameRestore)
		eventRestores, _ := p.eventRestoreFolder(action)
		for _, restore := range eventRestores {
			logCopy := proto.Clone(logBase).(*__.EventLog)
			logCopy.Event = &__.Event{Restore: restore}
			logAll = append(logAll, logCopy)
		}
	default:
		return logAll
	}
	return logAll
}

func (p *EventPlugin) ParseTagEvent(clientID string, event *common.SecurityLevelRequest) {
	millis := time.Now().UnixMilli()
	var app *__.Application
	app = &__.Application{
		Url:            wrapperspb.String(""),
		DisplayName:    wrapperspb.String("360"),
		ExecutableName: wrapperspb.String("360"),
		SecurityLevel:  wrapperspb.String("123"),
		ExeSha1:        wrapperspb.String("123"),
	}
	myLog := p.FillLog(app, common.EventInfoEventNameWriteTag, "allow", millis, millis)
	eventTag, err := p.eventWriteTag(event, millis)
	if err != nil {
		common.Logger.Errorf("写tag事件错误 err: %v", err)
		return
	}
	myLog.Event = &__.Event{}
	myLog.Event.WriteTag = eventTag
	marshal, err := proto.Marshal(myLog)
	if err != nil {
		return
	}
	asyncMsg := &common.Message{
		SourceID:    p.id,
		TargetID:    "tunnel",
		Destination: common.MessageSendLog,
		Type:        common.MessageInfo,
		Payload:     marshal,
	}
	// 创建消息总线
	bus := messagebus.GetMessageBusInstance()
	bus.SubmitAsyncTask(asyncMsg)
}

// Start 启动插件
func (p *EventPlugin) FillLog(application *__.Application, eventName, action string, eventTime, logTime int64) *__.EventLog {
	client := common.GetClientInstance()
	var eventLog __.EventLog
	logId, _ := common.DMRandomUuid()
	eventLog.LogId = wrapperspb.String(logId)
	eventLog.LogTime = wrapperspb.Int64(logTime)

	eventLog.User = wrapperspb.String(client.CurrentUser.Username)
	eventLog.Os = client.ClientOS
	eventLog.Application = application
	eventLog.Client = client.CurrentClient
	eventLog.Action = wrapperspb.String(action)
	eventLog.EventName = wrapperspb.String(eventName)
	eventLog.EventTime = wrapperspb.UInt64(uint64(eventTime))
	return &eventLog
}
