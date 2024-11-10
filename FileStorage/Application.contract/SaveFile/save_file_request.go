package contract

type SaveFileRequest struct {
	FileID   string `json:"file_id"`
	OwnerID  string `json:"owner_id"`
	FileName string `json:"file_name"`
	Content  []byte `json:"content"`
}
