package commands

import (
	"context"
	"identity-api/Application/contract"
	"identity-api/Domain/event"
	"identity-api/Domain/responses"
)

type RegisterUserCommandHandler struct {
	eventPublisher event.EventPublisher
}

func NewRegisterUserCommandHandler(eventPublisher event.EventPublisher) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{eventPublisher: eventPublisher}
}

func createResponse(statusCode int, message string, userID string) *contract.RegisterUserResponse {
	return &contract.RegisterUserResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		UserID:       userID,
	}
}

func (h *RegisterUserCommandHandler) Handle(ctx context.Context, command *RegisterUserCommand) (*contract.RegisterUserResponse, error) {
	// Logic to register user and publish event
	userID := "generated_user_id" // Replace with actual user ID generation logic
	return createResponse(201, "User registered successfully", userID), nil
}
