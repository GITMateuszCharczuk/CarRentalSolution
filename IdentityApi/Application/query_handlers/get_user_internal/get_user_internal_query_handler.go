package queries

import (
	"context"
	contract "identity-api/Application.contract/get_user_internal"
	"identity-api/Application/services"
	models "identity-api/Domain/models/user"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
)

type GetUserInternalQueryHandler struct {
	userQueryRepository repository_interfaces.UserQueryRepository
	tokenService        service_interfaces.JWTTokenService
}

func NewGetUserInternalQueryHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) *GetUserInternalQueryHandler {
	return &GetUserInternalQueryHandler{
		userQueryRepository: userQueryRepository,
		tokenService:        tokenService,
	}
}

func (h *GetUserInternalQueryHandler) Handle(ctx context.Context, query *GetUserInternalQuery) (*contract.GetUserInternalResponse, error) { //TODO: Restrtict from outside
	userID, _, err := h.tokenService.ValidateToken(query.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.GetUserInternalResponse](401, "Unauthorized")
		return &response, nil
	}

	user, err := h.userQueryRepository.GetUserByID(userID)
	if err != nil || user == nil {
		response := responses.NewResponse[contract.GetUserInternalResponse](404, "User not found")
		return &response, nil
	}
	userInfo := models.UserInfo{
		Name:         user.Name,
		Surname:      user.Surname,
		PhoneNumber:  user.PhoneNumber,
		EmailAddress: user.EmailAddress,
		Address:      user.Address,
		PostalCode:   user.PostalCode,
		City:         user.City,
	}

	return &contract.GetUserInternalResponse{
		BaseResponse: responses.NewBaseResponse(200, "User retrieved successfully"),
		UserSecureInfo: models.UserSecureInfo{
			UserInfo: userInfo,
			ID:       userID,
			Roles:    services.ConvertRolesToString(user.Roles),
		},
	}, nil
}
