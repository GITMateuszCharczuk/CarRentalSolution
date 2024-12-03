package commands

import (
	commands "email-service/Application/commmand_handlers/send_email"
	"email-service/Domain/event"
	"email-service/Infrastructure/config"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerSendEmailCommandHandler(eventPublisher event.EventPublisher, config *config.Config) {
	handler := commands.NewSendEmailCommandHandler(eventPublisher, config.DefaultEmailSender)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterCommandHandlers(eventPublisher event.EventPublisher, config *config.Config) {
	registerSendEmailCommandHandler(eventPublisher, config)
}
