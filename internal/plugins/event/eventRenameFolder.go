package event

import (
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (p *EventPlugin) eventRenameFolder(event *model.EventRenameFolder) (*__.RenameFolder, error) {
	var eventRenameFolder *__.RenameFolder
	eventRenameFolder = &__.RenameFolder{
		SrcFolder: wrapperspb.String(event.GetSrcFolder()),
		DesFolder: wrapperspb.String(event.GetDesFolder()),
	}
	return eventRenameFolder, nil
}
