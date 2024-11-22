package datafetcher

import (
	"email-service/Domain/models"
	mappers "email-service/Infrastructure/data_fetcher/mappers"
	responses "email-service/Infrastructure/data_fetcher/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

type DataFetcherImpl struct {
	mailHogUrl string
}

func NewDataFetcherImpl(mailHogUrl string) *DataFetcherImpl {
	return &DataFetcherImpl{mailHogUrl: mailHogUrl}
}

func (cmd *DataFetcherImpl) GetEmails() (*[]models.Email, error) {
	route := fmt.Sprintf("http://%s/api/v2/messages", cmd.mailHogUrl)
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
