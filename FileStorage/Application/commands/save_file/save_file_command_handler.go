// commands/save_file.go
package commands

import (
	"context"
	contract "file-storage/Application.contract/SaveFile"
	"file-storage/Domain/constants"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"fmt"
	"io"
	"path/filepath"

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

func (cmd *SaveFileCommandHandler) Handle(ctx context.Context, command *SaveFileCommand) (*contract.SaveFileResponse, error) {
	fileExtension := filepath.Ext(command.File.Filename)
	if fileExtension == "" {
		return &contract.SaveFileResponse{
			Title:   "StatusBadRequest",
			Message: "File must have a valid extension",
		}, nil
	}

	if !constants.IsAllowedExtension(fileExtension) {
		return &contract.SaveFileResponse{
			Title:   "StatusBadRequest",
			Message: fmt.Sprintf("File extension '%s' is not allowed", fileExtension),
		}, nil
	}

	u, err := uuid.NewUUID()
	if err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: "Error generating GUID",
		}, nil
	}

	fileId := u.String()

	fileContent, err := command.File.Open()
	if err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to open file: %v", err),
		}, nil
	}
	defer fileContent.Close()

	content, err := io.ReadAll(fileContent)
	if err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to read file content: %v", err),
		}, nil
	}

	fileData := models.File{
		ID:       fileId,
		OwnerID:  command.OwnerID,
		FileName: command.File.Filename,
		Content:  content,
	}

	if err := cmd.eventPublisher.PublishEvent("file-events.upload", fileData, models.EventTypeUpload); err != nil {
		return &contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to save file: %v", err),
		}, nil
	}

	return &contract.SaveFileResponse{
		Title:   "StatusOK",
		Message: "File saved successfully",
		Id:      fileId,
	}, nil
}
