// queries/get_file.go
package queries

import (
	"context"
	contract "file-storage/Application.contract/GetFile"
	"file-storage/Domain/repository_interfaces"
)

type GetFileQuery struct {
	Request  contract.GetFileRequest
	fileRepo repository_interfaces.FileRepository
}

func NewGetFileQuery(fileRepo repository_interfaces.FileRepository) *GetFileQuery {
	return &GetFileQuery{
		fileRepo: fileRepo,
	}
}

func (query *GetFileQuery) Execute() (contract.GetFileResponse, error) {
	file, err := query.fileRepo.GetFileByID(context.Background(), query.Request.FileID)
	if err != nil {
		return contract.GetFileResponse{
			Title:   "Error",
			Message: "File not found",
		}, err
	}

	return contract.GetFileResponse{
		Title:   "Success",
		Message: "File retrieved successfully",
		File:    file,
	}, nil
}
