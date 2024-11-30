package contract

import "file-storage/Domain/models"

type GetFileResponse struct {
	Title         string             `json:"title"`
	Message       string             `json:"message"`
	FileStream    *models.FileStream `json:"file_stream,omitempty"`
	FileExtension *string            `json:"file_extension,omitempty"`
}

type GetFileResponse404 struct {
	Title   string `json:"title" example:"Not Found" swaggertype:"string"`
	Message string `json:"message" example:"The requested file was not found." swaggertype:"string"`
}

type GetFileResponse200 struct {
	Title   string `json:"title" example:"StatusOK" swaggertype:"string"`
	Message string `json:"message" example:"File retrieved successfully." swaggertype:"string"`
}

type GetFileResponse400 struct {
	Title   string `json:"title" example:"Bad Request" swaggertype:"string"`
	Message string `json:"message" example:"Invalid file request." swaggertype:"string"`
}

type GetFileResponse500 struct {
	Title   string `json:"title" example:"Internal Server Error" swaggertype:"string"`
	Message string `json:"message" example:"An unexpected error occurred." swaggertype:"string"`
}
