package commands

import (
	send_email_command "file-storage/Application/commands/send_email"
	"file-storage/Domain/event"

	"github.com/google/wire"
)

func ProvideSaveFileCommandHandler(eventPublisher event.EventPublisher) *send_email_command.SendEmailCommandHandler {
	return send_email_command.NewSendEmailCommandHandler(eventPublisher)
}

type CommandHandlers struct {
	SendEmailCommand *send_email_command.SendEmailCommand
}

var WireSet = wire.NewSet(
	ProvideSaveFileCommandHandler,
	wire.Struct(new(CommandHandlers), "*"),
)
