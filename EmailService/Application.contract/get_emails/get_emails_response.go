package contract

import "email-service/Domain/models"

type GetEmailsResponse struct {
	Title   string         `json:"title" example:"Email List" swaggertype:"string"`
	Message string         `json:"message"`
	Emails  []models.Email `json:"emails"`
}

type GetEmailsResponse404 struct {
	Title   string         `json:"title" example:"Not Found" swaggertype:"string"`
	Message string         `json:"message" example:"No emails found." swaggertype:"string"`
	Emails  []models.Email `json:"emails" swaggertype:"array,object"`
}

type GetEmailsResponse200 struct {
	Title   string         `json:"title" example:"Success" swaggertype:"string"`
	Message string         `json:"message" example:"Emails retrieved successfully." swaggertype:"string"`
	Emails  []models.Email `json:"emails"`
}

type GetEmailsResponse400 struct {
	Title   string         `json:"title" example:"Bad Request" swaggertype:"string"`
	Message string         `json:"message" example:"Invalid request for emails." swaggertype:"string"`
	Emails  []models.Email `json:"emails" swaggertype:"array,object"`
}

type GetEmailsResponse500 struct {
	Title   string         `json:"title" example:"Internal Server Error" swaggertype:"string"`
	Message string         `json:"message" example:"An unexpected error occurred." swaggertype:"string"`
	Emails  []models.Email `json:"emails" swaggertype:"array,object"`
}
