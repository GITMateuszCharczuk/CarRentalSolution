// queries/get_file.go
package queries

import (
	"context"
	"file-storage/Domain/models"
	"file-storage/Domain/repository_interfaces"
)

type GetFileQuery struct {
	FileID   string
	OwnerID  string
	fileRepo repository_interfaces.FileRepository
}

func NewGetFileQuery(fileRepo repository_interfaces.FileRepository) *GetFileQuery {
	return &GetFileQuery{
		fileRepo: fileRepo,
	}
}

func (query *GetFileQuery) Execute() (models.File, error) {
	return query.fileRepo.GetFileByID(context.Background(), query.FileID)
}
