package datafetcher

import (
	"encoding/json"
	"fmt"
	"identity-api/Domain/models"
	pagination "identity-api/Domain/requests"
	mappers "identity-api/Infrastructure/data_fetcher/mappers"
	responses "identity-api/Infrastructure/data_fetcher/responses"
	"log"
	"net/http"
)

type DataFetcherImpl struct {
	mailHogUrl string
}

func NewDataFetcherImpl(mailHogUrl string) *DataFetcherImpl {
	return &DataFetcherImpl{mailHogUrl: mailHogUrl}
}

func (cmd *DataFetcherImpl) GetEmails(pagination pagination.Pagination) (*[]models.Email, error) {
	var route string
	log.Println("pagination", pagination)
	if pagination.Page > 0 && pagination.PageSize > 0 {
		offset := (pagination.Page - 1) * pagination.PageSize
		limit := pagination.PageSize
		log.Println("offset", offset)
		log.Println("limit", limit)
		route = fmt.Sprintf("http://%s/api/v2/messages?start=%d&limit=%d", cmd.mailHogUrl, offset, limit)
	} else {
		route = fmt.Sprintf("http://%s/api/v2/messages", cmd.mailHogUrl)
	}

	resp, err := http.Get(route)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve emails: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result responses.GetEmailsRawResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse email response: %w", err)
	}

	emails := mappers.MapMessagesToEmails(result.Messages)
	return &emails, nil
}
