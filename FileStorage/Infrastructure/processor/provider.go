package processor

import (
	"file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
)

func InitializeEventProcessor(fileRepo repository_interfaces.FileRepository) *EventProcessor {
	wire.Build(NewEventProcessor)
	return &EventProcessor{}
}

var WireSet = wire.NewSet(InitializeEventProcessor)
