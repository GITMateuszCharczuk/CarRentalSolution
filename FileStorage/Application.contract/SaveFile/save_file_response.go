package contract

type SaveFileResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Id      string `json:"id"`
}
