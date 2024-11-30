package commands

import (
	delete_command "file-storage/Application/commands/delete_file"
	save_command "file-storage/Application/commands/save_file"
	"file-storage/Domain/event"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerSaveFileCommandHandler(eventPublisher event.EventPublisher) {
	handler := save_command.NewSaveFileCommandHandler(eventPublisher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteFileCommandHandler(eventPublisher event.EventPublisher) {
	handler := delete_command.NewDeleteFileCommandHandler(eventPublisher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterCommandHandlers(eventPublisher event.EventPublisher) {
	registerSaveFileCommandHandler(eventPublisher)
	registerDeleteFileCommandHandler(eventPublisher)
}
