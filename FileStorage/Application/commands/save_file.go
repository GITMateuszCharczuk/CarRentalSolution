// commands/save_file.go
package commands

import (
	"context"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"file-storage/Domain/repository_interfaces"
)

type SaveFileCommand struct {
	FileID   string
	OwnerID  string
	FileName string
	Content  []byte

	fileRepo       repository_interfaces.FileRepository
	eventPublisher event.EventPublisher
}

func NewSaveFileCommand(fileRepo repository_interfaces.FileRepository, eventPublisher event.EventPublisher) *SaveFileCommand {
	return &SaveFileCommand{
		fileRepo:       fileRepo,
		eventPublisher: eventPublisher,
	}
}

func (cmd *SaveFileCommand) Execute() error {
	file := models.File{
		ID:       cmd.FileID,
		OwnerID:  cmd.OwnerID,
		FileName: cmd.FileName,
		Content:  cmd.Content,
	}

	if err := cmd.fileRepo.InsertFile(context.Background(), file); err != nil {
		return err
	}

	if err := cmd.eventPublisher.PublishEvent("FileSaved", file); err != nil {
		return err
	}

	return nil
}
