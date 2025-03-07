package event

import (
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
)

func (p *EventPlugin) eventSaveAs(event *model.EventSaveAs) (*__.CopyPast, error) {

	pathDes := event.GetDestinationFile()
	pathSource := event.GetSourceFile()

	return p.eventCopyTag(pathDes, pathSource)
}
