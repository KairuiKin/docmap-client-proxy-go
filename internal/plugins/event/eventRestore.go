package event

import (
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
)

func (p *EventPlugin) eventRestore(event *model.EventRestore) (*__.CopyPast, error) {
	pathDes := event.GetDestinationFile()
	pathSource := event.GetSourceFile()
	return p.eventMoveTag(pathDes, pathSource)
}
