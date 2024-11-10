// commands/delete_file.go
package commands

import (
	contract "file-storage/Application.contract/DeleteFile"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"file-storage/Domain/repository_interfaces"
	"fmt"
)

type DeleteFileCommand struct {
	Request        contract.DeleteFileRequest
	fileRepo       repository_interfaces.FileRepository
	eventPublisher event.EventPublisher
}

func NewDeleteFileCommand(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *DeleteFileCommand {
	return &DeleteFileCommand{
		fileRepo:       fileRepo,
		eventPublisher: eventPublisher,
	}
}

func (cmd *DeleteFileCommand) Execute() (contract.DeleteFileResponse, error) {
	if err := cmd.eventPublisher.PublishEvent("events.delete", cmd.Request.FileID, models.EventTypeDelete); err != nil {
		return contract.DeleteFileResponse{
			Title:   "Error",
			Message: fmt.Sprintf("Failed to delete file: %v", err),
		}, err
	}

	return contract.DeleteFileResponse{
		Title:   "Success",
		Message: "File deleted successfully",
	}, nil
}
