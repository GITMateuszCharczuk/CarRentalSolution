package contract

import "file-storage/Domain/models"

type GetFileResponse struct {
	Title   string      `json:"title"`
	Message string      `json:"message"`
	File    models.File `json:"file"`
}
