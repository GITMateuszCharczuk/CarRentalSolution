package contract

import "email-service/Domain/models"

type GetEmailResponse struct {
	Title   string       `json:"title"`
	Message string       `json:"message"`
	Email   models.Email `json:"email"`
}
