package contract

import (
	models "email-service/Domain/models/external"
	pagination "email-service/Domain/requests"
)

type GetEmailsRequest struct {
	pagination.Pagination `json:"pagination" binding:"required"`
	models.JwtToken       `json:"-"`
}
