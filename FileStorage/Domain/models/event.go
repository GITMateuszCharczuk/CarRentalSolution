// models/event.go

package models

type EventType string

const (
	EventTypeDelete EventType = "delete"
	EventTypeUpload EventType = "upload"
)

type Event struct {
	Type EventType   `json:"type"`
	Data interface{} `json:"data"`
}
