// queries/get_file.go
package queries

import (
	"context"
	contract "file-storage/Application.contract/GetFile"
	"file-storage/Domain/repository_interfaces"
	"log"
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

func (cmd *GetFileQueryHandler) Handle(ctx context.Context, query *GetFileQuery) (*contract.GetFileResponse, error) {
	if query.FileID == "" {
		return &contract.GetFileResponse{
			Title:   "StatusBadRequest",
			Message: "Missing required query parameters: file_id",
		}, nil
	}

	file, err := cmd.fileRepo.GetFileByID(ctx, query.FileID)
	if err != nil {
		log.Println("Error getting file:", err)
		return &contract.GetFileResponse{
			Title:   "StatusNotFound",
			Message: "File not found",
		}, nil
	}

	fileExtension := filepath.Ext(file.FileName)

	if fileExtension == "" {
		return &contract.GetFileResponse{
			Title:   "StatusBadRequest",
			Message: "File must have a valid extension",
		}, nil
	}

	return &contract.GetFileResponse{
		Title:         "StatusOK",
		Message:       "File retrieved successfully",
		FileStream:    &file,
		FileExtension: &fileExtension,
	}, nil
}
