package queries

import pagination "email-service/Domain/requests"

type GetEmailsQuery struct {
	pagination.Pagination
}
