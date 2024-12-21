package commands

import (
	"context"
	contract "email-service/Application.contract/send_email"
	"email-service/Application/utils"
	"email-service/Domain/constants"
	"email-service/Domain/event"
	"email-service/Domain/models"
	"email-service/Domain/responses"
	"email-service/Domain/service_interfaces"
	"fmt"
)

type SendEmailCommandHandler struct {
	eventPublisher        event.EventPublisher
	defaultEmailReceiver  string
	microserviceConnector service_interfaces.MicroserviceConnector
}

func NewSendEmailCommandHandler(eventPublisher event.EventPublisher, defaultEmailReceiver string, microserviceConnector service_interfaces.MicroserviceConnector) *SendEmailCommandHandler {
	return &SendEmailCommandHandler{
		eventPublisher:        eventPublisher,
		defaultEmailReceiver:  defaultEmailReceiver,
		microserviceConnector: microserviceConnector,
	}
}

func createResponse(statusCode int, message string) *contract.SendEmailResponse {
	return &contract.SendEmailResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
	}
}

func (cmd *SendEmailCommandHandler) Handle(ctx context.Context, command *SendEmailCommand) (*contract.SendEmailResponse, error) {
	tokenInfo, err := cmd.microserviceConnector.ValidateToken(command.JwtToken)
	if err != nil || !tokenInfo.Valid {
		return createResponse(401, "Invalid JWT token"), nil
	}
	if !utils.IsRole(constants.User, tokenInfo.Roles) {
		return createResponse(403, "Forbidden oparation"), nil
	}

	event := models.SendEmailEvent{
		From:    command.From,
		To:      cmd.defaultEmailReceiver,
		Subject: command.Subject,
		Body:    command.Body,
	}

	if command.From == cmd.defaultEmailReceiver {
		return createResponse(400, "Forbidden oparation"), nil
	}

	if err := cmd.eventPublisher.PublishEvent("email-events.send_email", event, models.EventTypeSendEmail); err != nil {
		return createResponse(500, fmt.Sprintf("Failed to send email: %v", err)), nil
	}

	return createResponse(200, "Email sent successfully"), nil
}
