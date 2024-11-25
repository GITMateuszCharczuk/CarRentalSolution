package processor

type EventProcessor interface {
	ProcessUploadEvent(data interface{}) error
	ProcessDeleteEvent(data interface{}) error
}
