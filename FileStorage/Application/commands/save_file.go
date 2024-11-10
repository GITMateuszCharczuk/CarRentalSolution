// commands/save_file.go
package commands

import (
	contract "file-storage/Application.contract/SaveFile"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"file-storage/Domain/repository_interfaces"
	"fmt"
)

type SaveFileCommand struct {
	Request        contract.SaveFileRequest
	fileRepo       repository_interfaces.FileRepository
	eventPublisher event.EventPublisher
}

func NewSaveFileCommand(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *SaveFileCommand {
	return &SaveFileCommand{
		fileRepo:       fileRepo,
		eventPublisher: eventPublisher,
	}
}

func (cmd *SaveFileCommand) Execute() (contract.SaveFileResponse, error) {
	fileData := models.File{
		ID:       cmd.Request.FileID,
		OwnerID:  cmd.Request.OwnerID,
		FileName: cmd.Request.FileName,
		Content:  cmd.Request.Content,
	}

	if err := cmd.eventPublisher.PublishEvent("events.upload", fileData, models.EventTypeUpload); err != nil {
		return contract.SaveFileResponse{
			Title:   "Error",
			Message: fmt.Sprintf("Failed to save file: %v", err),
		}, err
	}

	return contract.SaveFileResponse{
		Title:   "Success",
		Message: "File saved successfully",
	}, nil
}
