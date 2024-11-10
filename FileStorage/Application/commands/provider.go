package commands

import (
	delete_command "file-storage/Application/commands/delete_file"
	save_command "file-storage/Application/commands/save_file"
	"file-storage/Domain/event"
	"file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
)

func ProvideSaveFileCommandHandler(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *save_command.SaveFileCommandHandler {
	return save_command.NewSaveFileCommandHandler(fileRepo, eventPublisher)
}

func ProvideDeleteFileCommandHandler(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *delete_command.DeleteFileCommandHandler {
	return delete_command.NewDeleteFileCommandHandler(fileRepo, eventPublisher)
}

type CommandHandlers struct {
	SaveFileCommand   *save_command.SaveFileCommand
	DeleteFileCommand *delete_command.DeleteFileCommand
}

var WireSet = wire.NewSet(
	ProvideSaveFileCommandHandler,
	ProvideDeleteFileCommandHandler,
	wire.Struct(new(CommandHandlers), "*"),
)
