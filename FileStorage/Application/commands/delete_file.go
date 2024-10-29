// commands/delete_file.go
package commands

import (
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"file-storage/Domain/repository_interfaces"
	"fmt"
)

type DeleteFileCommand struct {
	FileID  string
	OwnerID string

	fileRepo       repository_interfaces.FileRepository
	eventPublisher event.EventPublisher
}

func NewDeleteFileCommand(fileRepo repository_interfaces.FileRepository) *DeleteFileCommand {
	return &DeleteFileCommand{
		fileRepo: fileRepo,
	}
}

func (cmd *DeleteFileCommand) Execute() error {
	// if err := cmd.fileRepo.DeleteFileByID(context.Background(), cmd.FileID); err != nil {
	// 	return err
	// }

	if err := cmd.eventPublisher.PublishEvent("events.delete", cmd.FileID, models.EventTypeDelete); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}
	return nil
}
