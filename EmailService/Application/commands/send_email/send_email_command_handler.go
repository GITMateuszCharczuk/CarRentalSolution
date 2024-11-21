package commands

import (
	contract "email-service/Application.contract/send_email"
	"email-service/Domain/event"
	"email-service/Domain/models"
	"fmt"
)

type SendEmailCommandHandler struct {
	eventPublisher event.EventPublisher
}

func NewSendEmailCommandHandler(eventPublisher event.EventPublisher) *SendEmailCommandHandler {
	return &SendEmailCommandHandler{
		eventPublisher: eventPublisher,
	}
}

func (cmd *SendEmailCommandHandler) Execute(command SendEmailCommand) (contract.SendEmailResponse, error) {
	event := models.SendEmailEvent{
		From:    command.From,
		To:      command.To,
		Subject: command.Subject,
		Body:    command.Body,
	}
	if err := cmd.eventPublisher.PublishEvent("events.send_email", event, models.EventTypeSend); err != nil {
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
