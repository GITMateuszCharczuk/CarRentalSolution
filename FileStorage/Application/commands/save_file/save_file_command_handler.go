// commands/save_file.go
package commands

import (
	contract "file-storage/Application.contract/SaveFile"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"file-storage/Domain/repository_interfaces"
	"fmt"
)

type SaveFileCommandHandler struct {
	fileRepo       repository_interfaces.FileRepository
	eventPublisher event.EventPublisher
}

func NewSaveFileCommandHandler(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *SaveFileCommandHandler {
	return &SaveFileCommandHandler{
		fileRepo:       fileRepo,
		eventPublisher: eventPublisher,
	}
}

func (cmd *SaveFileCommandHandler) Execute(command SaveFileCommand) (*contract.SaveFileResponse, error) {
	fileData := models.File{
		ID:       command.FileID,
		OwnerID:  command.OwnerID,
		FileName: command.FileName,
		Content:  command.Content,
	}

	if err := cmd.eventPublisher.PublishEvent("events.upload", fileData, models.EventTypeUpload); err != nil {
		return &contract.SaveFileResponse{
			Title:   "Error",
			Message: fmt.Sprintf("Failed to save file: %v", err),
		}, err
	}

	return &contract.SaveFileResponse{
		Title:   "Success",
		Message: "File saved successfully",
	}, nil
}
