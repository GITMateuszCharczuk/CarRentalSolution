package processor

type EventProcessor interface {
	ProcessSendEmailEvent(data interface{}) error
}
