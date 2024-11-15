package contract

import "mime/multipart"

type SaveFileRequest struct {
	OwnerID string                `json:"owner_id"`
	File    *multipart.FileHeader `json:"file"`
}
