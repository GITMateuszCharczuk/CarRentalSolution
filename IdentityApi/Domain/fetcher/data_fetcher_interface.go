package fetcher

import "identity-api/Domain/models"
import pagination "identity-api/Domain/requests"

type DataFetcher interface {
	GetEmails(pagination pagination.Pagination) (*[]models.Email, error)
}
