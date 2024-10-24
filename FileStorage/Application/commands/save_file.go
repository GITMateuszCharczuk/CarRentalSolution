package commands

import (
    "context"
    "file-storage/Domain/models"
    "file-storage/Infrastructure/db"
    "file-storage/Infrastructure/queue"
)

type SaveFileCommand struct {
    FileID   string
    OwnerID  string
    FilePath string
}

func SaveFile(cmd SaveFileCommand) error {
    file := models.File{
        ID:       cmd.FileID,
        OwnerID:  cmd.OwnerID,
        FilePath: cmd.FilePath,
    }
    // Save file to MongoDB
    if err := db.InsertFile(context.Background(), file.ID, file.OwnerID, file.FilePath); err != nil {
        return err
    }

    // Publish event to NATS JetStream
    queue.PublishEvent("FileSaved")

    return nil
}
