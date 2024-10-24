package commands

import (
    "context"
    "file-storage/Infrastructure/db"
)

type DeleteFileCommand struct {
    FileID  string
    OwnerID string
}

func DeleteFile(cmd DeleteFileCommand) error {
    // Logic for deleting a file in MongoDB
    if err := db.DeleteFileByID(context.Background(), cmd.FileID); err != nil {
        return err
    }
    return nil
}
