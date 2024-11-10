package commands

import (
	"file-storage/Domain/event"
	"file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
)

func ProvideSaveFileCommand(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *SaveFileCommand {
	return NewSaveFileCommand(fileRepo, eventPublisher)
}

func ProvideDeleteFileCommand(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *DeleteFileCommand {
	return NewDeleteFileCommand(fileRepo, eventPublisher)
}

type Commands struct {
	SaveFileCommand   *SaveFileCommand
	DeleteFileCommand *DeleteFileCommand
}

var WireSet = wire.NewSet(
	ProvideSaveFileCommand,
	ProvideDeleteFileCommand,
	wire.Struct(new(Commands), "*"),
)
