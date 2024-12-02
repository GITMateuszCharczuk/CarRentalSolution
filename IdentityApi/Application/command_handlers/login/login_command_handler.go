package commands

import (
	"context"
	"identity-api/Application/contract"
	"identity-api/Domain/responses"
)

type LoginCommandHandler struct{}

func NewLoginCommandHandler() *LoginCommandHandler {
	return &LoginCommandHandler{}
}

func createResponse(statusCode int, message string, token string, roles []string) *contract.LoginResponse {
	return &contract.LoginResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		Token:        token,
		Roles:        roles,
	}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, command *LoginCommand) (*contract.LoginResponse, error) {
	// Logic to authenticate user and generate token
	token := "generated_jwt_token" // Replace with actual token generation logic
	roles := []string{"user"}      // Replace with actual role retrieval logic
	return createResponse(200, "Login successful", token, roles), nil
}
