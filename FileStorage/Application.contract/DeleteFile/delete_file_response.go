package contract

type DeleteFileResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type DeleteFileResponse404 struct {
	Title   string `json:"title" example:"Not Found" swaggertype:"string"`
	Message string `json:"message" example:"The requested file was not found." swaggertype:"string"`
}

type DeleteFileResponse200 struct {
	Title   string `json:"title" example:"StatusOK" swaggertype:"string"`
	Message string `json:"message" example:"File deleted successfully." swaggertype:"string"`
}

type DeleteFileResponse400 struct {
	Title   string `json:"title" example:"Bad Request" swaggertype:"string"`
	Message string `json:"message" example:"Invalid delete file request." swaggertype:"string"`
}

type DeleteFileResponse500 struct {
	Title   string `json:"title" example:"Internal Server Error" swaggertype:"string"`
	Message string `json:"message" example:"An unexpected error occurred." swaggertype:"string"`
}
