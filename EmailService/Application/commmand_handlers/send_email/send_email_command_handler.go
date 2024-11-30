package commands

import (
	"context"
	contract "email-service/Application.contract/send_email"
	"email-service/Domain/event"
	"email-service/Domain/models"
	"email-service/Domain/responses"
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

func createResponse(statusCode int, message string) *contract.SendEmailResponse {
	return &contract.SendEmailResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
	}
}

func (cmd *SendEmailCommandHandler) Handle(ctx context.Context, command *SendEmailCommand) (*contract.SendEmailResponse, error) {
	event := models.SendEmailEvent{
		From:    command.From,
		To:      command.To,
		Subject: command.Subject,
		Body:    command.Body,
	}

	if !(command.From == cmd.defaultEmailSender || command.To == cmd.defaultEmailSender) {
		return createResponse(400, "From or to is not the desired email address"), nil
	}

	if err := cmd.eventPublisher.PublishEvent("email-events.send_email", event, models.EventTypeSendEmail); err != nil {
		return createResponse(500, fmt.Sprintf("Failed to send email: %v", err)), nil
	}

	return createResponse(200, "Email sent successfully"), nil
}
