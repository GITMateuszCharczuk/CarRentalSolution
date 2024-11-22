package processor

import (
	smtp "email-service/Infrastructure/email_sender"
	mappers "email-service/Infrastructure/event_processor/mappers"
	"log"
)

type EventProcessorImpl struct {
	sender smtp.EmailSender
}

func NewEventProcessor(sender smtp.EmailSender) *EventProcessorImpl {
	return &EventProcessorImpl{sender: sender}
}

func (p *EventProcessorImpl) ProcessSendEmailEvent(data interface{}) error {
	email, err := mappers.MapToSendEmailEvent(data)
	if err != nil {
		return err
	}

	if err := p.sender.SendEmail(email.From, email.To, email.Subject, email.Body); err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Printf("Successfully processed 'send_email' event for email to: %s", email.To)
	return nil
}
