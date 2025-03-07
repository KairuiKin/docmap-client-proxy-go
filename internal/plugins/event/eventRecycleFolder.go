package event

import (
	"docmap-client-proxy-go/internal/common"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"io/fs"
	"path/filepath"
	"strings"
)

func (p *EventPlugin) eventRecycleFolder(event *model.EventRecycleFolder) ([]*__.CopyPast, error) {
	pathDes := event.GetDesFolder()
	pathSource := event.GetSrcFolder()
	var eventMoves []*__.CopyPast
	// 使用 filepath.WalkDir 递归遍历目录
	err := filepath.WalkDir(pathDes, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			common.Logger.Errorf("无法访问 %q: %v\n", path, err)
			return nil
		}

		// 如果是文件，打印文件路径
		if !d.IsDir() {
			dstFile := path
			srcFile := strings.Replace(dstFile, pathDes, pathSource, 1)
			eventMove, err := p.eventMoveTag(dstFile, srcFile)
			if err == nil {
				eventMoves = append(eventMoves, eventMove)
			}

		}

		return nil
	})

	if err != nil {
		common.Logger.Errorf("遍历目录时出错: %v\n", err)
	}

	return eventMoves, nil
}
