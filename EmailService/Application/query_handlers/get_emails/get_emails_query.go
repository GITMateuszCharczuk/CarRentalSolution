package queries

import (
	models "email-service/Domain/models/external"
	pagination "email-service/Domain/requests"
)

type GetEmailsQuery struct {
	pagination.Pagination
	JwtToken models.JwtToken
}
