package event

import "file-storage/Domain/models"

type EventPublisher interface {
	PublishEvent(subject string, data interface{}, eventType models.EventType) error
}
