package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"encoding/json"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (p *EventPlugin) eventWriteTag(event *common.SecurityLevelRequest, logTime int64) (*__.WriteTag, error) {
	var eventWriteTag *__.WriteTag
	eventWriteTag = &__.WriteTag{}
	fileId := &common.FileIdentity{}
	//通过同步的方式获取fileID以及密级信息
	// 创建消息总线
	bus := messagebus.GetMessageBusInstance()
	{
		for _, cfg := range p.config.Classification.AllowedSecret {
			if cfg.Name == event.SecurityLevel {
				newSecret := &common.DocumentClassification{
					FilePath: event.FilePath,
					FileLevel: common.Secret{
						Level: cfg.Level,
						Code:  cfg.Code,
						Name:  cfg.Name,
					},
				}
				asyncMsg := &common.Message{
					SourceID:    p.id,
					TargetID:    "tag",
					Destination: common.MessageSetFileLevel,
					Type:        common.MessageInfo,
					Payload:     newSecret,
				}
				rep, err := bus.Send(asyncMsg)
				if err != nil {
					return nil, err
				}
				fileLevel := rep.Payload.(*common.DocumentClassification)
				secretData, _ := json.Marshal(fileLevel.FileLevel)
				oldFileTag := &__.FileTagEntry{
					Key:   wrapperspb.String("secret"),
					Value: wrapperspb.String(string(secretData)),
					Time:  wrapperspb.UInt64(uint64(logTime)),
				}
				eventWriteTag.BeforeFileTags = append(eventWriteTag.BeforeFileTags, oldFileTag)
				fileLevel = rep.Payload.(*common.DocumentClassification)
				secretData, _ = json.Marshal(newSecret.FileLevel)
				fileTag := &__.FileTagEntry{
					Key:   wrapperspb.String("secret"),
					Value: wrapperspb.String(string(secretData)),
					Time:  wrapperspb.UInt64(uint64(logTime)),
				}
				eventWriteTag.AfterFileTags = append(eventWriteTag.AfterFileTags, fileTag)
				//获取文件tag
				asyncMsg = &common.Message{
					SourceID:    p.id,
					TargetID:    "tag",
					Destination: common.MessageGetFileID,
					Type:        common.MessageInfo,
					Payload:     &event.FilePath,
				}
				rep, err = bus.Send(asyncMsg)
				if err != nil {
					return nil, err
				}
				fileId = rep.Payload.(*common.FileIdentity)
				family := &__.FileTagEntry{
					Key:   wrapperspb.String("family"),
					Value: wrapperspb.String(fileId.Family),
					Time:  wrapperspb.UInt64(uint64(logTime)),
				}
				eventWriteTag.BeforeFileTags = append(eventWriteTag.BeforeFileTags, family)
				eventWriteTag.AfterFileTags = append(eventWriteTag.AfterFileTags, family)
				nextFileId := &__.FileTagEntry{
					Key:   wrapperspb.String("identity"),
					Value: wrapperspb.String(fileId.Identity),
					Time:  wrapperspb.UInt64(uint64(logTime)),
				}
				eventWriteTag.BeforeFileTags = append(eventWriteTag.BeforeFileTags, nextFileId)
				eventWriteTag.AfterFileTags = append(eventWriteTag.AfterFileTags, nextFileId)
				break
			}
		}
	}
	eventWriteTag.SourceFile, _ = GetFileInfo(event.FilePath, fileId.Identity)
	eventWriteTag.DestinationFile = eventWriteTag.SourceFile
	return eventWriteTag, nil
}
