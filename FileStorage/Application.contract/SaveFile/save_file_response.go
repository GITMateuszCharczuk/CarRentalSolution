package contract

type SaveFileResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Id      string `json:"id,omitempty"`
}

type SaveFileResponse404 struct {
	Title   string `json:"title" example:"Not Found" swaggertype:"string"`
	Message string `json:"message" example:"The requested file was not found." swaggertype:"string"`
	Id      string `json:"id" swaggertype:"string"`
}

type SaveFileResponse201 struct {
	Title   string `json:"title" example:"StatusCreated" swaggertype:"string"`
	Message string `json:"message" example:"File saved successfully." swaggertype:"string"`
	Id      string `json:"id"`
}

type SaveFileResponse400 struct {
	Title   string `json:"title" example:"Bad Request" swaggertype:"string"`
	Message string `json:"message" example:"Invalid save file request." swaggertype:"string"`
	Id      string `json:"id" swaggertype:"string"`
}

type SaveFileResponse500 struct {
	Title   string `json:"title" example:"Internal Server Error" swaggertype:"string"`
	Message string `json:"message" example:"An unexpected error occurred." swaggertype:"string"`
	Id      string `json:"id" swaggertype:"string"`
}
