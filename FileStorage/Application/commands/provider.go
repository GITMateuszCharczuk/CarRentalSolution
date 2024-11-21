package commands

import (
	delete_command "file-storage/Application/commands/delete_file"
	save_command "file-storage/Application/commands/save_file"
	"file-storage/Domain/event"

	"github.com/google/wire"
)

func ProvideSaveFileCommandHandler(eventPublisher event.EventPublisher) *save_command.SaveFileCommandHandler {
	return save_command.NewSaveFileCommandHandler(eventPublisher)
}

func ProvideDeleteFileCommandHandler(eventPublisher event.EventPublisher) *delete_command.DeleteFileCommandHandler {
	return delete_command.NewDeleteFileCommandHandler(eventPublisher)
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
