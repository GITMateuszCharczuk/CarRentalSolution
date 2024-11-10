package contract

type GetFileResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	File    []byte `json:"file"`
}
