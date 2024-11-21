package processor

import (
	"email-service/Infrastructure/smtp"

	"github.com/google/wire"
)

func InitializeEventProcessor(sender smtp.EmailSender) *EventProcessor {
	return NewEventProcessor(sender)
}

var WireSet = wire.NewSet(NewEventProcessor)
