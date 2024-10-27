// commands/commands.go

package commands

import (
	"file-storage/Domain/event"
	"file-storage/Domain/repository"

	"github.com/google/wire"
)

type Commands struct {
	SaveFile   *SaveFileCommand
	DeleteFile *DeleteFileCommand
}

func ProvideCommands(fileRepo repository.FileRepository, eventPublisher event.EventPublisher) *Commands {
	return &Commands{
		SaveFile:   NewSaveFileCommand(fileRepo, eventPublisher),
		DeleteFile: NewDeleteFileCommand(fileRepo),
	}
}

var WireSet = wire.NewSet(ProvideCommands)
