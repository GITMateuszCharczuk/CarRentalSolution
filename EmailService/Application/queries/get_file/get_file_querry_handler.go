// queries/get_file.go
package queries

import (
	"context"
	contract "email-service/Application.contract/GetFile"
	"email-service/Domain/models"
	"email-service/Domain/repository_interfaces"
	"path/filepath"
)

type GetFileQueryHandler struct {
	fileRepo repository_interfaces.FileRepository
}

func NewGetFileQueryHandler(fileRepo repository_interfaces.FileRepository) *GetFileQueryHandler {
	return &GetFileQueryHandler{
		fileRepo: fileRepo,
	}
}

func (cmd *GetFileQueryHandler) Execute(query GetFileQuery) (contract.GetFileResponse, *models.FileStream, *string, error) {
	file, err := cmd.fileRepo.GetFileByID(context.Background(), query.FileID)
	if err != nil {
		return contract.GetFileResponse{
			Title:   "StatusNotFound",
			Message: "File not found",
		}, nil, nil, err
	}

	fileExtension := filepath.Ext(file.FileName)
	if fileExtension == "" {
		return contract.GetFileResponse{
			Title:   "StatusBadRequest",
			Message: "File must have a valid extension",
		}, nil, nil, err
	}
	resp := contract.GetFileResponse{
		Title:   "StatusOK",
		Message: "File retrieved successfully",
	}

	return resp, &file, &fileExtension, nil
}
