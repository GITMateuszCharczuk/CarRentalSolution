package commands

import models "file-storage/Domain/models/external"

type DeleteFileCommand struct {
	FileID          string
	models.JwtToken `json:",inline"`
}
