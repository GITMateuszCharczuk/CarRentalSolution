package queries

import (
	"context"
	contract "identity-api/Application.contract/get_user_id"
	"identity-api/Application/services"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
)

type GetUserIDQueryHandler struct {
	userQueryRepository repository_interfaces.UserQueryRepository
	tokenService        service_interfaces.JWTTokenService
}

func NewGetUserIDQueryHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) *GetUserIDQueryHandler {
	return &GetUserIDQueryHandler{
		userQueryRepository: userQueryRepository,
		tokenService:        tokenService,
	}
}

func (h *GetUserIDQueryHandler) Handle(ctx context.Context, query *GetUserIDQuery) (*contract.GetUserIDResponse, error) { //TODO: Restrtict from outside
	userID, _, err := h.tokenService.ValidateToken(query.JwtToken)
	if err != nil {
		return &contract.GetUserIDResponse{
			BaseResponse: responses.NewBaseResponse(401, "Unauthorized"),
		}, nil
	}

	user, err := h.userQueryRepository.GetUserByID(userID)
	if err != nil || user == nil {
		return &contract.GetUserIDResponse{
			BaseResponse: responses.NewBaseResponse(404, "User not found"),
		}, nil
	}

	return &contract.GetUserIDResponse{
		BaseResponse: responses.NewBaseResponse(200, "User retrieved successfully"),
		UserID:       userID,
		Roles:        services.ConvertRolesToString(user.Roles),
	}, nil
}
