// commands/delete_file.go
package commands

import (
	contract "email-service/Application.contract/DeleteFile"
	"email-service/Domain/event"
	"email-service/Domain/models"
	"email-service/Domain/repository_interfaces"
	"fmt"
)

type DeleteFileCommandHandler struct {
	fileRepo       repository_interfaces.FileRepository
	eventPublisher event.EventPublisher
}

func NewDeleteFileCommandHandler(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *DeleteFileCommandHandler {
	return &DeleteFileCommandHandler{
		fileRepo:       fileRepo,
		eventPublisher: eventPublisher,
	}
}

func (cmd *DeleteFileCommandHandler) Execute(command DeleteFileCommand) (contract.DeleteFileResponse, error) {
	if err := cmd.eventPublisher.PublishEvent("events.delete", command.FileID, models.EventTypeDelete); err != nil {
		return contract.DeleteFileResponse{ //TODO dodać checka
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to delete file: %v", err),
		}, err
	}

	return contract.DeleteFileResponse{
		Title:   "StatusOK",
		Message: "File deleted successfully",
	}, nil
}
