// queries/get_file.go
package queries

import (
	"context"
	contract "file-storage/Application.contract/GetFile"
	"file-storage/Domain/repository_interfaces"
)

type GetFileQueryHandler struct {
	fileRepo repository_interfaces.FileRepository
}

func NewGetFileQueryHandler(fileRepo repository_interfaces.FileRepository) *GetFileQueryHandler {
	return &GetFileQueryHandler{
		fileRepo: fileRepo,
	}
}

func (cmd *GetFileQueryHandler) Execute(query GetFileQuery) (contract.GetFileResponse, error) {
	file, err := cmd.fileRepo.GetFileByID(context.Background(), query.FileID)
	if err != nil {
		return contract.GetFileResponse{
			Title:   "StatusNotFound",
			Message: "File not found",
		}, err
	}

	return contract.GetFileResponse{
		Title:   "StatusOK",
		Message: "File retrieved successfully",
		File:    file,
	}, nil
}
