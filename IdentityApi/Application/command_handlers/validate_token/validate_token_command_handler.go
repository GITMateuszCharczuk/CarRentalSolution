package commands

import (
	"context"
	"identity-api/Application/contract"
	"identity-api/Domain/responses"
)

type ValidateTokenCommandHandler struct{}

func NewValidateTokenCommandHandler() *ValidateTokenCommandHandler {
	return &ValidateTokenCommandHandler{}
}

func createResponse(statusCode int, message string, roles []string) *contract.ValidateTokenResponse {
	return &contract.ValidateTokenResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		Roles:        roles,
	}
}

func (h *ValidateTokenCommandHandler) Handle(ctx context.Context, command *ValidateTokenCommand) (*contract.ValidateTokenResponse, error) {
	// Logic to validate token
	roles := []string{"user"} // Replace with actual role retrieval logic
	return createResponse(200, "Token is valid", roles), nil
}
