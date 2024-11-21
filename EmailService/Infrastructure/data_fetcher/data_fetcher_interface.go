package datafetcher

import "email-service/Domain/models"

type DataFetcher interface {
	GetEmails() (*[]models.Email, error)
}
