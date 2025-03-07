package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (p *EventPlugin) eventMove(event *model.EventMove) (*__.CopyPast, error) {
	pathDes := event.GetDestinationFile()
	pathSource := event.GetSourceFile()
	return p.eventMoveTag(pathDes, pathSource)
}

func (p *EventPlugin) eventMoveTag(pathDes, pathSource string) (*__.CopyPast, error) {
	if pathDes == pathSource {
		//拷贝事件因为来源文件和目标文件是同一个被忽略！
		return nil, nil
	}
	fileIdNew := &common.IdentityWriteTask{}
	bus := messagebus.GetMessageBusInstance()
	fileTagSet := &common.SetIdentityRequest{
		FilePath:       pathDes,
		ParentFilePath: pathSource,
		IsMove:         1,
	}
	//设置目标文件Tag
	asyncMsg := &common.Message{
		SourceID:    p.id,
		TargetID:    "tag",
		Destination: common.MessageSetFileID,
		Type:        common.MessageInfo,
		Payload:     fileTagSet,
	}
	rep, err := bus.Send(asyncMsg)
	if err != nil {
		common.Logger.Warnf("%v", err)
	} else {
		fileIdNew = rep.Payload.(*common.IdentityWriteTask)
	}
	//获取原文件密级
	asyncMsg = &common.Message{
		SourceID:    p.id,
		TargetID:    "tag",
		Destination: common.MessageGetFileLevel,
		Type:        common.MessageInfo,
		Payload:     &pathSource,
	}
	rep, err = bus.Send(asyncMsg)
	if err != nil {

	} else {
		//删除文件tag相关信息
		asyncMsg = &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageDeleteFile,
			Type:        common.MessageInfo,
			Payload:     &pathSource,
		}
		bus.Send(asyncMsg)
		//获取到了 则设置目标级别，并且删除原文件tag和级别
		fileLevel := rep.Payload.(*common.DocumentClassification)
		//设置目标文件密级
		newDC := &common.DocumentClassification{
			FilePath:  pathDes,
			FileLevel: fileLevel.FileLevel,
		}
		// 获取FileLevel
		asyncMsg = &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageSetFileLevel,
			Type:        common.MessageInfo,
			Payload:     newDC,
		}
		bus.Send(asyncMsg)
	}

	var eventCopy *__.CopyPast
	eventCopy = &__.CopyPast{}

	eventCopy.DestinationFile, _ = GetFileInfo(pathDes, fileIdNew.Identity.Identity)
	eventCopy.SourceFile = &__.FileInfo{
		EntityIdentity:     wrapperspb.String(fileIdNew.Identity.Identity),
		ContentIdentities:  eventCopy.DestinationFile.ContentIdentities,
		Path:               wrapperspb.String(pathSource),
		Filename:           eventCopy.DestinationFile.Filename,
		DriveType:          eventCopy.DestinationFile.DriveType,
		SecurityLevel:      eventCopy.DestinationFile.SecurityLevel,
		ClassificationType: eventCopy.DestinationFile.ClassificationType,
		CreateTime:         eventCopy.DestinationFile.CreateTime,
		ModifyTime:         eventCopy.DestinationFile.ModifyTime,
		Sha1:               eventCopy.DestinationFile.Sha1,
		Size:               eventCopy.DestinationFile.Size,
		Ext:                eventCopy.DestinationFile.Ext,
		EncryptChannel:     eventCopy.DestinationFile.EncryptChannel,
		Encrypted:          eventCopy.DestinationFile.Encrypted,
	}
	return eventCopy, nil
}
