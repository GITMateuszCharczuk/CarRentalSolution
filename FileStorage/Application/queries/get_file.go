package queries

import (
    "context"
    "file-storage/Infrastructure/db"
)

type GetFileQuery struct {
    FileID  string
    OwnerID string
}

func GetFile(query GetFileQuery) (string, error) {
    // Get file from MongoDB by ID
    return db.GetFileByID(context.Background(), query.FileID)
}
