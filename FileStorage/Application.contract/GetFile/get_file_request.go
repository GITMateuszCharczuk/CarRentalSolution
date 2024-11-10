package contract

type GetFileRequest struct {
	FileID  string `json:"file_id"`
	OwnerID string `json:"owner_id"`
}
