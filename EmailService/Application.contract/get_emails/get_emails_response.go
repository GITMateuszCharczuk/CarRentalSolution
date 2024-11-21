package contract

import "email-service/Domain/models"

type GetEmailsResponse struct {
	Title   string         `json:"title"`
	Message string         `json:"message"`
	Emails  []models.Email `json:"emails"`
}
