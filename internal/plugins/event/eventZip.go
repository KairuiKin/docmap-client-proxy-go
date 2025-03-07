package event

import (
	"docmap-client-proxy-go/internal/common"
	"docmap-client-proxy-go/internal/messagebus"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"fmt"
	"strconv"
)

func (p *EventPlugin) eventZip(event *model.EventZip) (*__.Zip, error) {

	var SourceFiles []*__.FileZipMap
	var DesFiles []*__.FileInfo

	SourceFile := event.GetSourceFiles()
	DesFile := event.GetDestinationFiles()
	fileLevelLast := &common.DocumentClassification{
		FilePath: "",
		FileLevel: common.Secret{
			Level: "0",
			Code:  "",
			Name:  "",
		},
	}
	bus := messagebus.GetMessageBusInstance()
	for _, zipMap := range SourceFile {
		//获取文件tag
		asyncMsg := &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageGetFileID,
			Type:        common.MessageInfo,
			Payload:     &zipMap.Src.File,
		}
		rep, err := bus.Send(asyncMsg)
		if err == nil {
			fileId := rep.Payload.(*common.FileIdentity)
			srcFile, _ := GetFileInfo(zipMap.Src.File, fileId.Identity)
			SourceFiles = append(SourceFiles, &__.FileZipMap{
				Src: srcFile,
				Des: nil,
			})
		}
		//获取原文件密级
		asyncMsg = &common.Message{
			SourceID:    p.id,
			TargetID:    "tag",
			Destination: common.MessageGetFileLevel,
			Type:        common.MessageInfo,
			Payload:     &zipMap.Src.File,
		}
		rep, err = bus.Send(asyncMsg)
		if err != nil {
			return nil, err
		}
		fileLevel := rep.Payload.(*common.DocumentClassification)
		// 将字符串转换为整数
		n1, err1 := strconv.Atoi(fileLevel.FileLevel.Level)
		if err1 != nil {
			fmt.Println("转换 s1 失败：", err1)
			continue
		}

		n2, err2 := strconv.Atoi(fileLevelLast.FileLevel.Level)
		if err2 != nil {
			fmt.Println("转换 s2 失败：", err2)
			continue
		}
		if n1 > n2 {
			fileLevelLast = fileLevel
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
			srcFile, _ := GetFileInfo(file.File, fileId.Identity)
			DesFiles = append(DesFiles, srcFile)
			//设置文件密级
			//设置目标文件密级
			newDC := &common.DocumentClassification{
				FilePath:  file.File,
				FileLevel: fileLevelLast.FileLevel,
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
	}
	zip := &__.Zip{
		SourceFiles:      SourceFiles,
		DestinationFiles: DesFiles,
	}
	return zip, nil
}
