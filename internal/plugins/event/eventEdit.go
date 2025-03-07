package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (p *EventPlugin) eventEdit(event *model.EventEdit) (*__.CopyPast, error) {
	filePath := event.GetDestinationFile()
	//获取文件tag
	asyncMsg := &common.Message{
		SourceID:    p.id,
		TargetID:    "tag",
		Destination: common.MessageGetFileID,
		Type:        common.MessageInfo,
		Payload:     &filePath,
	}
	// 创建消息总线
	fileId := &common.FileIdentity{}
	bus := messagebus.GetMessageBusInstance()
	rep, err := bus.Send(asyncMsg)
	if err != nil {

	} else {
		fileId = rep.Payload.(*common.FileIdentity)
	}
	var eventEdit *__.CopyPast
	eventEdit = &__.CopyPast{}
	eventEdit.SourceFile, _ = GetFileInfo(filePath, fileId.Identity)
	eventEdit.SourceFile.Sha1 = wrapperspb.String("")
	eventEdit.DestinationFile, _ = GetFileInfo(filePath, fileId.Identity)
	return nil, nil
}
