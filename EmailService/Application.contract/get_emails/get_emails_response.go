package contract

import "file-storage/Domain/models"

type GetEmailsResponse struct {
	Title   string         `json:"title"`
	Message string         `json:"message"`
	Emails  []models.Email `json:"emails"`
}
