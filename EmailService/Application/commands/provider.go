package commands

import (
	send_email_command "email-service/Application/commands/send_email"
	"email-service/Domain/event"
	"email-service/Infrastructure/config"
	"github.com/google/wire"
)

func ProvideSendEmailCommandHandler(eventPublisher event.EventPublisher, config *config.Config) *send_email_command.SendEmailCommandHandler {
	return send_email_command.NewSendEmailCommandHandler(eventPublisher, config.DefaultEmailSender)
}

type CommandHandlers struct {
	SendEmailCommand *send_email_command.SendEmailCommandHandler
}

var WireSet = wire.NewSet(
	ProvideSendEmailCommandHandler,
	wire.Struct(new(CommandHandlers), "*"),
)
