package commands

import (
	models "file-storage/Domain/models/external"
	"mime/multipart"
)

type SaveFileCommand struct {
	models.JwtToken `json:",inline"`
	File            *multipart.FileHeader `json:"file"`
}
