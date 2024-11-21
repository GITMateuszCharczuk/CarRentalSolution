package processor

import (
	"file-storage/Infrastructure/smtp"

	"github.com/google/wire"
)

func InitializeEventProcessor(sender smtp.EmailSender) *EventProcessor {
	return NewEventProcessor(sender)
}

var WireSet = wire.NewSet(NewEventProcessor)
