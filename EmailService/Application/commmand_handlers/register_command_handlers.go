package commands

import (
	commands "email-service/Application/commmand_handlers/send_email"
	send_internal_email_commands "email-service/Application/commmand_handlers/send_internal_email"
	"email-service/Domain/event"
	"email-service/Domain/service_interfaces"
	"email-service/Infrastructure/config"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerSendEmailCommandHandler(eventPublisher event.EventPublisher, config *config.Config, microserviceConnector service_interfaces.MicroserviceConnector) {
	handler := commands.NewSendEmailCommandHandler(eventPublisher, config.DefaultEmailSender, microserviceConnector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerSendInternalEmailCommandHandler(eventPublisher event.EventPublisher, config *config.Config) {
	handler := send_internal_email_commands.NewSendInternalEmailCommandHandler(eventPublisher, config.DefaultEmailSender)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterCommandHandlers(eventPublisher event.EventPublisher, config *config.Config, microserviceConnector service_interfaces.MicroserviceConnector) {
	registerSendEmailCommandHandler(eventPublisher, config, microserviceConnector)
	registerSendInternalEmailCommandHandler(eventPublisher, config)
}
