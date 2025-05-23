package repository_interfaces

import (
	"context"
	"file-storage/Domain/models"
)

type FileRepository interface {
	InsertFile(ctx context.Context, file models.File) error
	GetFileByID(ctx context.Context, fileID string) (models.FileStream, error)
	DeleteFileByID(ctx context.Context, fileID string) error
}
