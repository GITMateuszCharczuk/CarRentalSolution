package event

import "github.com/nats-io/nats.go"

type EventReceiver interface {
	ProcessEvent(m *nats.Msg)
	StartReceiving() error
	StopReceiving() error
}
