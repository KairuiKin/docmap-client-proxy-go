package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (p *EventPlugin) eventDelete(event *model.EventDelete) (*__.Delete, error) {

	fileId := &common.FileIdentity{}
	//获取文件tag
	asyncMsg := &common.Message{
		SourceID:    p.id,
		TargetID:    "tag",
		Destination: common.MessageGetFileID,
		Type:        common.MessageInfo,
		Payload:     &event.Path,
	}
	// 创建消息总线
	bus := messagebus.GetMessageBusInstance()
	rep, err := bus.Send(asyncMsg)
	if err != nil {
	} else {
		fileId = rep.Payload.(*common.FileIdentity)
		if fileId == nil {
			fileId = &common.FileIdentity{}
		}
	}
	var eventDelete *__.Delete
	eventDelete = &__.Delete{}
	eventDelete.SourceFile = &__.FileInfo{
		EntityIdentity:     wrapperspb.String(fileId.Identity),
		Path:               wrapperspb.String(event.Path),
		Filename:           wrapperspb.String(event.Filename),
		DriveType:          wrapperspb.String(event.DriveType),
		SecurityLevel:      wrapperspb.String(event.SecurityLevel),
		CreateTime:         wrapperspb.UInt64(event.CreateTime),
		ModifyTime:         wrapperspb.UInt64(event.ModifyTime),
		Sha1:               wrapperspb.String(event.Sha1),
		Size:               wrapperspb.UInt64(event.Size),
		Ext:                wrapperspb.String(event.Ext),
		EncryptChannel:     wrapperspb.String(""), // 设置为空字符串
		Encrypted:          wrapperspb.Int32(0),   // 设置为 false (0)
		ClassificationType: make([]*wrapperspb.StringValue, 0),
	}
	for _, s := range event.ClassificationType {
		eventDelete.SourceFile.ClassificationType = append(eventDelete.SourceFile.ClassificationType, wrapperspb.String(s))
	}
	//获取文件tag
	asyncMsg = &common.Message{
		SourceID:    p.id,
		TargetID:    "tag",
		Destination: common.MessageDeleteFile,
		Type:        common.MessageInfo,
		Payload:     &event.Path,
	}
	bus.Send(asyncMsg)
	return eventDelete, nil
}
