// models/event.go

package models

type EventType string

const (
	EventTypeSend EventType = "send"
)

type Event struct {
	Type EventType   `json:"type"`
	Data interface{} `json:"data"`
}
