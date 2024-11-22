package processor

import (
	"email-service/Infrastructure/email_sender"

	"github.com/google/wire"
)

func InitializeEventProcessor(sender smtp.EmailSender) *EventProcessorImpl {
	return NewEventProcessor(sender)
}

var WireSet = wire.NewSet(NewEventProcessor)
