package event

type EventPublisher interface {
	PublishEvent(subject string, data interface{}) error
}
