// commands/delete_file.go
package commands

import (
	contract "file-storage/Application.contract/DeleteFile"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"fmt"
)

type DeleteFileCommandHandler struct {
	eventPublisher event.EventPublisher
}

func NewDeleteFileCommandHandler(eventPublisher event.EventPublisher) *DeleteFileCommandHandler {
	return &DeleteFileCommandHandler{
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
