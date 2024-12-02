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
	Success bool         `json:"success" example:"false" swaggertype:"boolean"`
	Message string       `json:"message" example:"The requested email was not found." swaggertype:"string"`
	Email   models.Email `json:"email" swaggertype:"object"`
}

type GetEmailResponse200 struct {
	Success bool         `json:"success" example:"true" swaggertype:"boolean"`
	Message string       `json:"message" example:"Email retrieved successfully." swaggertype:"string"`
	Email   models.Email `json:"email"`
}

type GetEmailResponse400 struct {
	Success bool         `json:"success" example:"false" swaggertype:"boolean"`
	Message string       `json:"message" example:"Invalid email request." swaggertype:"string"`
	Email   models.Email `json:"email" swaggertype:"object"`
}

type GetEmailResponse500 struct {
	Success bool         `json:"success" example:"false" swaggertype:"boolean"`
	Message string       `json:"message" example:"An unexpected error occurred." swaggertype:"string"`
	Email   models.Email `json:"email" swaggertype:"object"`
}
