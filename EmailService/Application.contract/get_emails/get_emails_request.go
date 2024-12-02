package contract

import pagination "email-service/Domain/requests"

type GetEmailsRequest struct {
	pagination.Pagination `json:"pagination" binding:"required"`
}
