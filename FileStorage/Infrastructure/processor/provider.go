package processor

import (
	"file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
)

func InitializeEventProcessor(fileRepo repository_interfaces.FileRepository) EventProcessor {
	return NewEventProcessorImpl(fileRepo)
}

var WireSet = wire.NewSet(NewEventProcessorImpl)
