package contract

import (
	"email-service/Domain/models"
	"email-service/Domain/responses"
)

type GetEmailResponse struct {
	responses.BaseResponse
	Email models.Email `json:"email,omitempty"`
}

type GetEmailResponse404 struct {
	Title   string       `json:"title" example:"Not Found" swaggertype:"string"`
	Message string       `json:"message" example:"The requested email was not found." swaggertype:"string"`
	Email   models.Email `json:"email" swaggertype:"object"` // Empty email object
}

type GetEmailResponse200 struct {
	Title   string       `json:"title" example:"Success" swaggertype:"string"`
	Message string       `json:"message" example:"Email retrieved successfully." swaggertype:"string"`
	Email   models.Email `json:"email"`
}

type GetEmailResponse400 struct {
	Title   string       `json:"title" example:"Bad Request" swaggertype:"string"`
	Message string       `json:"message" example:"Invalid email request." swaggertype:"string"`
	Email   models.Email `json:"email" swaggertype:"object"` // Empty email object
}

type GetEmailResponse500 struct {
	Title   string       `json:"title" example:"Internal Server Error" swaggertype:"string"`
	Message string       `json:"message" example:"An unexpected error occurred." swaggertype:"string"`
	Email   models.Email `json:"email" swaggertype:"object"` // Empty email object
}
