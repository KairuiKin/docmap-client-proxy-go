package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
)

func (p *EventPlugin) eventCopy(event *model.EventCopy) (*__.CopyPast, error) {
	pathDes := event.GetDestinationFile()
	pathSource := event.GetSourceFile()
	return p.eventCopyTag(pathDes, pathSource)
}

func (p *EventPlugin) eventCopyTag(pathDes, pathSource string) (*__.CopyPast, error) {
	if pathDes == pathSource {
		//拷贝事件因为来源文件和目标文件是同一个被忽略！
		return nil, nil
	}
	fileId := &common.FileIdentity{}
	fileIdNew := &common.IdentityWriteTask{}
	//获取文件tag
	asyncMsg := &common.Message{
		SourceID:    p.id,
		TargetID:    "tag",
		Destination: common.MessageGetFileID,
		Type:        common.MessageInfo,
		Payload:     &pathSource,
	}
	// 创建消息总线
	bus := messagebus.GetMessageBusInstance()
	rep, err := bus.Send(asyncMsg)
	if err != nil {
		common.Logger.Warnf("%v", err)
	} else {
		fileId = rep.Payload.(*common.FileIdentity)
		fileTagSet := &common.SetIdentityRequest{
			FilePath:       pathDes,
			ParentFilePath: pathSource,
			IsMove:         0,
		}
		//设置目标文件Tag
		asyncMsg = &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageSetFileID,
			Type:        common.MessageInfo,
			Payload:     fileTagSet,
		}
		rep, err = bus.Send(asyncMsg)
		if err != nil {
			common.Logger.Warnf("%v", err)
		}
		fileIdNew = rep.Payload.(*common.IdentityWriteTask)
		//获取原文件密级
		asyncMsg = &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageGetFileLevel,
			Type:        common.MessageInfo,
			Payload:     &pathSource,
		}
		rep, err = bus.Send(asyncMsg)
		if err == nil {
			fileLevel := rep.Payload.(*common.DocumentClassification)
			//设置目标文件密级
			newDC := &common.DocumentClassification{
				FilePath:  pathDes,
				FileLevel: fileLevel.FileLevel,
			}
			asyncMsg = &common.Message{
				SourceID:    p.id,
				TargetID:    "tag",
				Destination: common.MessageSetFileLevel,
				Type:        common.MessageInfo,
				Payload:     newDC,
			}
			rep, err = bus.Send(asyncMsg)
			if err != nil {
				return nil, err
			}
		}
	}
	var eventCopy *__.CopyPast
	eventCopy = &__.CopyPast{}
	eventCopy.SourceFile, _ = GetFileInfo(pathSource, fileId.Identity)
	eventCopy.DestinationFile, _ = GetFileInfo(pathDes, fileIdNew.Identity.Identity)
	return eventCopy, nil
}
