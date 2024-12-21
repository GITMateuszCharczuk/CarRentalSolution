package contract

import "email-service/Domain/responses"

type SendInternalEmailResponse struct {
	responses.BaseResponse
}

type SendInternalEmailResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Email sent successfully." swaggertype:"string"`
}

type SendInternalEmailResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request for sending email." swaggertype:"string"`
}

type SendInternalEmailResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"An unexpected error occurred while sending email." swaggertype:"string"`
}
