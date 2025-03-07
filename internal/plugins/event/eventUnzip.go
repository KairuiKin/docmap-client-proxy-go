package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
)

func (p *EventPlugin) eventUnzip(event *model.EventUnzip) (*__.Unzip, error) {
	var SourceFiles []*__.FileInfo
	var DesFiles []*__.FileUnzipMap
	SourceFile := event.GetSourceFiles()
	DesFile := event.GetDestinationFiles()
	// 创建消息总线
	bus := messagebus.GetMessageBusInstance()
	for _, file := range SourceFile {
		//获取文件tag
		asyncMsg := &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageGetFileID,
			Type:        common.MessageInfo,
			Payload:     &file.File,
		}
		rep, err := bus.Send(asyncMsg)
		if err == nil {
			fileId := rep.Payload.(*common.FileIdentity)
			srcFile, _ := GetFileInfo(file.File, fileId.Identity)
			SourceFiles = append(SourceFiles, srcFile)
		}
	}
	for _, file := range DesFile {
		//获取文件tag
		asyncMsg := &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageGetFileID,
			Type:        common.MessageInfo,
			Payload:     &file.File,
		}
		rep, err := bus.Send(asyncMsg)
		if err == nil {
			fileId := rep.Payload.(*common.FileIdentity)
			desFile, _ := GetFileInfo(file.File, fileId.Identity)
			DesFiles = append(DesFiles, &__.FileUnzipMap{
				Src: &__.FileInZipInfo{
					EntityIdentity: nil,
					ZipPath:        nil,
					Size:           nil,
					Sha1:           nil,
					InZipTime:      nil,
					Encrypted:      nil,
				},
				Des: desFile,
			})
		}
	}
	unZip := &__.Unzip{
		SourceFiles:      SourceFiles,
		DestinationFiles: DesFiles,
	}
	return unZip, nil
}
