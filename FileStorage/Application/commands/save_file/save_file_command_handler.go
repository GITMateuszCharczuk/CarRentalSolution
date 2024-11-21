// commands/save_file.go
package commands

import (
	contract "file-storage/Application.contract/SaveFile"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"fmt"
	"io"

	"github.com/google/uuid"
)

type SaveFileCommandHandler struct {
	eventPublisher event.EventPublisher
}

func NewSaveFileCommandHandler(eventPublisher event.EventPublisher) *SaveFileCommandHandler {
	return &SaveFileCommandHandler{
		eventPublisher: eventPublisher,
	}
}

func (cmd *SaveFileCommandHandler) Execute(command SaveFileCommand) (*contract.SaveFileResponse, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: "Error generating GUID",
		}, err
	}
	fileId := u.String()

	fileContent, err := command.File.Open()
	if err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to open file: %v", err),
		}, err
	}
	defer fileContent.Close()

	content, err := io.ReadAll(fileContent)
	if err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to read file content: %v", err),
		}, err
	}

	fileData := models.File{
		ID:       fileId,
		OwnerID:  command.OwnerID,
		FileName: command.File.Filename,
		Content:  content,
	}

	if err := cmd.eventPublisher.PublishEvent("events.upload", fileData, models.EventTypeUpload); err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to save file: %v", err),
		}, err
	}

	return &contract.SaveFileResponse{
		Title:   "StatusOK",
		Message: "File saved successfully",
		Id:      fileId,
	}, nil
}
