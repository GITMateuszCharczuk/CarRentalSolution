package commands

import "mime/multipart"

type SaveFileCommand struct {
	OwnerID string                `json:"owner_id"`
	File    *multipart.FileHeader `json:"file"`
}
