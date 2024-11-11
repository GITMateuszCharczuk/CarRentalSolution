package contract

type SaveFileRequest struct {
	OwnerID  string `json:"owner_id"`
	FileName string `json:"file_name"`
	Content  []byte `json:"content"`
}
