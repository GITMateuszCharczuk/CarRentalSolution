package event

import "email-service/Domain/models"

type EventPublisher interface {
	PublishEvent(subject string, data interface{}, eventType models.EventType) error
}
