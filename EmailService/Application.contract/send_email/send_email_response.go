package contract

import "email-service/Domain/responses"

type SendEmailResponse struct {
	responses.BaseResponse
}

type SendEmailResponse200 struct {
	Title   string `json:"title" example:"Success" swaggertype:"string"`
	Message string `json:"message" example:"Email sent successfully." swaggertype:"string"`
}

type SendEmailResponse400 struct {
	Title   string `json:"title" example:"Bad Request" swaggertype:"string"`
	Message string `json:"message" example:"Invalid request for sending email." swaggertype:"string"`
}

type SendEmailResponse500 struct {
	Title   string `json:"title" example:"Internal Server Error" swaggertype:"string"`
	Message string `json:"message" example:"An unexpected error occurred while sending email." swaggertype:"string"`
}
