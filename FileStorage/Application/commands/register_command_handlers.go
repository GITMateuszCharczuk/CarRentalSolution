package commands

import (
	delete_command "file-storage/Application/commands/delete_file"
	save_command "file-storage/Application/commands/save_file"
	"file-storage/Domain/event"
	interfaces "file-storage/Domain/service_interfaces"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerSaveFileCommandHandler(eventPublisher event.EventPublisher, microserviceConnector interfaces.MicroserviceConnector) {
	handler := save_command.NewSaveFileCommandHandler(eventPublisher, microserviceConnector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteFileCommandHandler(eventPublisher event.EventPublisher, microserviceConnector interfaces.MicroserviceConnector) {
	handler := delete_command.NewDeleteFileCommandHandler(eventPublisher, microserviceConnector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterCommandHandlers(eventPublisher event.EventPublisher, microserviceConnector interfaces.MicroserviceConnector) {
	registerSaveFileCommandHandler(eventPublisher, microserviceConnector)
	registerDeleteFileCommandHandler(eventPublisher, microserviceConnector)
}
