package datafetcher

import (
	"email-service/Domain/models"
	responses "email-service/Infrastructure/data_fetcher/responses"
	utils "email-service/Infrastructure/data_fetcher/utils"
)

func mapRawResponseToEmail(message *responses.GetEmailsResponse) models.Email {
	return models.Email{
		ID:      message.ID,
		From:    message.From.Mailbox + "@" + message.From.Domain,
		To:      utils.FormatToAddresses(message.To),
		Subject: utils.ExtractFirstString(message.Content.Headers.Subject),
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
