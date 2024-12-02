package contract

import "email-service/Domain/responses"

type SendEmailResponse struct {
	responses.BaseResponse
}

type SendEmailResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Email sent successfully." swaggertype:"string"`
}

type SendEmailResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request for sending email." swaggertype:"string"`
}

type SendEmailResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"An unexpected error occurred while sending email." swaggertype:"string"`
}
