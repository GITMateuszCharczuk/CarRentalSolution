package contract

type DeleteFileRequest struct {
	FileID  string `json:"file_id"`
	OwnerID string `json:"owner_id"`
}
