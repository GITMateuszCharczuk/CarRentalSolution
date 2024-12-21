package contract

import models "file-storage/Domain/models/external"

type DeleteFileRequest struct {
	FileID          string `json:"file_id"`
	models.JwtToken `json:",inline"`
}
