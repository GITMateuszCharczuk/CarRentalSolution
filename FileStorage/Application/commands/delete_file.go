// commands/delete_file.go
package commands

import (
	"context"
	"file-storage/Domain/repository_interfaces"
)

type DeleteFileCommand struct {
	fileRepo repository_interfaces.FileRepository
	FileID   string
	OwnerID  string
}

func NewDeleteFileCommand(fileRepo repository_interfaces.FileRepository) *DeleteFileCommand {
	return &DeleteFileCommand{
		fileRepo: fileRepo,
	}
}

func (cmd *DeleteFileCommand) Execute() error {
	if err := cmd.fileRepo.DeleteFileByID(context.Background(), cmd.FileID); err != nil {
		return err
	}
	return nil
}
