package commands

import (
	"context"
	contract "email-service/Application.contract/send_internal_email"
	"email-service/Domain/event"
	"email-service/Domain/models"
	"email-service/Domain/responses"
	"fmt"
)

type SendInternalEmailCommandHandler struct {
	eventPublisher     event.EventPublisher
	defaultEmailSender string
}

func NewSendInternalEmailCommandHandler(eventPublisher event.EventPublisher, defaultEmailSender string) *SendInternalEmailCommandHandler {
	return &SendInternalEmailCommandHandler{
		eventPublisher:     eventPublisher,
		defaultEmailSender: defaultEmailSender,
	}
}

func createResponse(statusCode int, message string) *contract.SendInternalEmailResponse {
	return &contract.SendInternalEmailResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
	}
}

func (cmd *SendInternalEmailCommandHandler) Handle(ctx context.Context, command *SendInternalEmailCommand) (*contract.SendInternalEmailResponse, error) {
	event := models.SendEmailEvent{
		From:    cmd.defaultEmailSender,
		To:      command.To,
		Subject: command.Subject,
		Body:    command.Body,
	}

	if command.To == cmd.defaultEmailSender {
		return createResponse(400, "Forbidden oparation"), nil
	}

	if err := cmd.eventPublisher.PublishEvent("email-events.send_email", event, models.EventTypeSendEmail); err != nil {
		return createResponse(500, fmt.Sprintf("Failed to send email: %v", err)), nil
	}

	return createResponse(200, "Email sent successfully"), nil
}
