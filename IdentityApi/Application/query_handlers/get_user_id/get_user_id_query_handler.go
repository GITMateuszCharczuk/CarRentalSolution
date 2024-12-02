package queries

import (
	"context"
	"identity-api/Application/contract"
	"identity-api/Domain/responses"
)

type GetUserIDQueryHandler struct{}

func NewGetUserIDQueryHandler() *GetUserIDQueryHandler {
	return &GetUserIDQueryHandler{}
}

func createResponse(statusCode int, message string, userID string) *contract.GetUserIDResponse {
	return &contract.GetUserIDResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		UserID:       userID,
	}
}

func (h *GetUserIDQueryHandler) Handle(ctx context.Context, query *GetUserIDQuery) (*contract.GetUserIDResponse, error) {
	// Logic to retrieve user ID based on token
	userID := "retrieved_user_id" // Replace with actual user ID retrieval logic
	return createResponse(200, "User ID retrieved successfully", userID), nil
}
