package fetcher

import "email-service/Domain/models"
import pagination "email-service/Domain/requests"

type DataFetcher interface {
	GetEmails(pagination pagination.Pagination) (*[]models.Email, error)
}
