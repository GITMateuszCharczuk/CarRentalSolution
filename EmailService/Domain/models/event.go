// models/event.go

package models

type EventType string

const (
	EventTypeSendEmail EventType = "send_email"
)

type Event struct {
	Type EventType   `json:"type"`
	Data interface{} `json:"data"`
}
