package contract

import (
	models "file-storage/Domain/models/external"
	"mime/multipart"
)

type SaveFileRequest struct {
	models.JwtToken `json:",inline"`
	File            *multipart.FileHeader `json:"file"`
}
