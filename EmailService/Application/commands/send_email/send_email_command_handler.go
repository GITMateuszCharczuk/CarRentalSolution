package commands

import (
	contract "email-service/Application.contract/send_email"
	"email-service/Domain/event"
	"email-service/Domain/models"
	"fmt"
)

type SendEmailCommandHandler struct {
	eventPublisher     event.EventPublisher
	defaultEmailSender string
}

func NewSendEmailCommandHandler(eventPublisher event.EventPublisher, defaultEmailSender string) *SendEmailCommandHandler {
	return &SendEmailCommandHandler{
		eventPublisher:     eventPublisher,
		defaultEmailSender: defaultEmailSender,
	}
}

func (cmd *SendEmailCommandHandler) Execute(command SendEmailCommand) (contract.SendEmailResponse, error) {
	event := models.SendEmailEvent{
		From:    command.From,
		To:      command.To,
		Subject: command.Subject,
		Body:    command.Body,
	}
	if !(command.From == cmd.defaultEmailSender || command.To == cmd.defaultEmailSender) {
		return contract.SendEmailResponse{
			Title:   "StatusBadRequest",
			Message: "From or to is not the desired email address",
		}, nil
	}
	if err := cmd.eventPublisher.PublishEvent("email-events.send_email", event, models.EventTypeSendEmail); err != nil {
		return contract.SendEmailResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to delete file: %v", err),
		}, err
	}

	return contract.SendEmailResponse{
		Title:   "StatusOK",
		Message: "Email sent successfully",
	}, nil
}
