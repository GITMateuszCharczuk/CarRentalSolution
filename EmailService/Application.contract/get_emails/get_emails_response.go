package contract

import (
	"email-service/Domain/models"
	"email-service/Domain/responses"
)

type GetEmailsResponse struct {
	responses.BaseResponse
	Emails []models.Email `json:"emails,omitempty"`
}

type GetEmailsResponse200 struct {
	Success bool           `json:"success" example:"true" swaggertype:"boolean"`
	Message string         `json:"message" example:"Emails retrieved successfully." swaggertype:"string"`
	Emails  []models.Email `json:"emails"`
}

type GetEmailsResponse404 struct {
	Success bool           `json:"success" example:"false" swaggertype:"boolean"`
	Message string         `json:"message" example:"No emails found." swaggertype:"string"`
	Emails  []models.Email `json:"emails" swaggertype:"array,object"`
}

type GetEmailsResponse400 struct {
	Success bool           `json:"success" example:"false" swaggertype:"boolean"`
	Message string         `json:"message" example:"Invalid request for emails." swaggertype:"string"`
	Emails  []models.Email `json:"emails" swaggertype:"array,object"`
}

type GetEmailsResponse500 struct {
	Success bool           `json:"success" example:"false" swaggertype:"boolean"`
	Message string         `json:"message" example:"An unexpected error occurred." swaggertype:"string"`
	Emails  []models.Email `json:"emails" swaggertype:"array,object"`
}
