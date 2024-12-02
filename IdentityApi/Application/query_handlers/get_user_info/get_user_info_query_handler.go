package queries

import (
	"context"
	"identity-api/Application/contract"
	"identity-api/Domain/models"
	"identity-api/Domain/responses"
)

type GetUserInfoQueryHandler struct{}

func NewGetUserInfoQueryHandler() *GetUserInfoQueryHandler {
	return &GetUserInfoQueryHandler{}
}

func createResponse(statusCode int, message string, userInfo models.UserSecureInfo) *contract.GetUserInfoResponse {
	return &contract.GetUserInfoResponse{
		BaseResponse:   responses.NewBaseResponse(statusCode, message),
		UserSecureInfo: userInfo,
	}
}

func (h *GetUserInfoQueryHandler) Handle(ctx context.Context, query *GetUserInfoQuery) (*contract.GetUserInfoResponse, error) {
	// Logic to retrieve user info based on token
	userInfo := models.UserSecureInfo{} // Replace with actual user info retrieval logic
	return createResponse(200, "User info retrieved successfully", userInfo), nil
}
