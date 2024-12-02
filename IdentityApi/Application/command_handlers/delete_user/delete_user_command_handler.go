package commands

import (
	"context"
	"identity-api/Application/contract"
	"identity-api/Domain/responses"
)

type DeleteUserCommandHandler struct{}

func NewDeleteUserCommandHandler() *DeleteUserCommandHandler {
	return &DeleteUserCommandHandler{}
}

func createResponse(statusCode int, message string) *contract.DeleteUserResponse {
	return &contract.DeleteUserResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
	}
}

func (h *DeleteUserCommandHandler) Handle(ctx context.Context, command *DeleteUserCommand) (*contract.DeleteUserResponse, error) {
	// Logic to delete user based on token
	return createResponse(200, "User deleted successfully"), nil
}
