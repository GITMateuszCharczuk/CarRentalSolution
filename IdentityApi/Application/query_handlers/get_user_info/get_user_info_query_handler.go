package queries

import (
	"context"
	contract "identity-api/Application.contract/get_user_info"
	"identity-api/Application/services"
	models "identity-api/Domain/models/user"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
)

type GetUserInfoQueryHandler struct {
	userQueryRepository repository_interfaces.UserQueryRepository
	tokenService        service_interfaces.JWTTokenService
}

func NewGetUserInfoQueryHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) *GetUserInfoQueryHandler {
	return &GetUserInfoQueryHandler{
		userQueryRepository: userQueryRepository,
		tokenService:        tokenService,
	}
}

func (h *GetUserInfoQueryHandler) Handle(ctx context.Context, query *GetUserInfoQuery) (*contract.GetUserInfoResponse, error) {
	requesterID, requesterRoles, err := h.tokenService.ValidateToken(query.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.GetUserInfoResponse](401, "Unauthorized")
		return &response, nil
	}

	isAdmin := services.IsAdminOrSuperAdmin(requesterRoles)

	userID := requesterID

	if query.Id != "" && isAdmin {
		userID = query.Id
	}

	existingUser, err := h.userQueryRepository.GetUserByID(userID)
	if err != nil || existingUser == nil {
		response := responses.NewResponse[contract.GetUserInfoResponse](404, "User not found")
		return &response, nil
	}

	userInfo := models.UserSecureInfo{
		UserInfo: models.UserInfo{
			Name:         existingUser.Name,
			Surname:      existingUser.Surname,
			PhoneNumber:  existingUser.PhoneNumber,
			EmailAddress: existingUser.EmailAddress,
			Address:      existingUser.Address,
			PostalCode:   existingUser.PostalCode,
			City:         existingUser.City,
		},
		Roles: services.ConvertRolesToString(existingUser.Roles),
	}

	if isAdmin {
		userInfo.ID = existingUser.ID
	}

	return &contract.GetUserInfoResponse{
		BaseResponse:   responses.NewBaseResponse(200, "User info retrieved successfully"),
		UserSecureInfo: userInfo,
	}, nil
}
