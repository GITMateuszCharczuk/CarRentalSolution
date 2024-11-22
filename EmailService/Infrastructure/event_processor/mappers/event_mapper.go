package processor

import (
	"email-service/Domain/models"
	"errors"
	"fmt"
)

var (
	ErrInvalidDataType = errors.New("invalid data type")
)

func MapToSendEmailEvent(data interface{}) (models.SendEmailEvent, error) {
	eventData, ok := data.(map[string]interface{})
	if !ok {
		return models.SendEmailEvent{}, fmt.Errorf("%w: expected map[string]interface{}", ErrInvalidDataType)
	}

	email := models.SendEmailEvent{}

	if from, ok := eventData["from"].(string); ok {
		email.From = from
	} else {
		return models.SendEmailEvent{}, fmt.Errorf("invalid type for 'from'")
	}

	if to, ok := eventData["to"].(string); ok {
		email.To = to
	} else {
		return models.SendEmailEvent{}, fmt.Errorf("invalid type for 'to'")
	}

	if subject, ok := eventData["subject"].(string); ok {
		email.Subject = subject
	} else {
		return models.SendEmailEvent{}, fmt.Errorf("invalid type for 'subject'")
	}

	if body, ok := eventData["body"].(string); ok {
		email.Body = body
	} else {
		return models.SendEmailEvent{}, fmt.Errorf("invalid type for 'body'")
	}

	return email, nil
}
