package queries

import (
	"context"
	"identity-api/Application/contract"
	"identity-api/Domain/models"
	"identity-api/Domain/responses"
)

type GetAllUsersQueryHandler struct{}

func NewGetAllUsersQueryHandler() *GetAllUsersQueryHandler {
	return &GetAllUsersQueryHandler{}
}

func createResponse(statusCode int, message string, users []models.UserSecureInfo) *contract.GetAllUsersResponse {
	return &contract.GetAllUsersResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		Users:        users,
	}
}

func (h *GetAllUsersQueryHandler) Handle(ctx context.Context, query *GetAllUsersQuery) (*contract.GetAllUsersResponse, error) {
	// Logic to retrieve all users
	users := []models.UserSecureInfo{} // Replace with actual user retrieval logic
	return createResponse(200, "Users retrieved successfully", users), nil
}
