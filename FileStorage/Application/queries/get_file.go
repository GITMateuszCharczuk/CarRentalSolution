// queries/get_file.go
package queries

import (
	"context"
	"file-storage/Domain/models"
	"file-storage/Domain/repository"
)

type GetFileQuery struct {
	FileID   string
	OwnerID  string
	fileRepo repository.FileRepository
}

func NewGetFileQuery(fileRepo repository.FileRepository) *GetFileQuery {
	return &GetFileQuery{
		fileRepo: fileRepo,
	}
}

func (query *GetFileQuery) Execute() (models.File, error) {
	return query.fileRepo.GetFileByID(context.Background(), query.FileID)
}
