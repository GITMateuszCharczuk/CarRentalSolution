package contract

import "file-storage/Domain/models"

type GetEmailResponse struct {
	Title   string       `json:"title"`
	Message string       `json:"message"`
	Email   models.Email `json:"email"`
}
