package datafetcher

import (
	"email-service/Domain/models"
	responses "email-service/Infrastructure/data_fetcher/responses"
)

func mapRawResponseToEmail(message *responses.GetEmailsResponse) models.Email {
	return models.Email{
		ID:      message.ID,
		From:    message.From,
		To:      message.To,
		Subject: message.Subject,
		Body:    message.Content.Body,
	}
}

func MapMessagesToEmails(messages []responses.GetEmailsResponse) []models.Email {
	emails := make([]models.Email, len(messages))
	for i, message := range messages {
		emails[i] = mapRawResponseToEmail(&message)
	}
	return emails
}
