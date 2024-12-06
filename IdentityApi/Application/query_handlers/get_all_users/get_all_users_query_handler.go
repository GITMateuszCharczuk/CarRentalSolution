package queries

import (
	"context"
	contract "identity-api/Application.contract/get_all_users"
	"identity-api/Application/services"
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/user"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
)

type GetAllUsersQueryHandler struct {
	userQueryRepository repository_interfaces.UserQueryRepository
	tokenService        service_interfaces.JWTTokenService
}

func NewGetAllUsersQueryHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) *GetAllUsersQueryHandler {
	return &GetAllUsersQueryHandler{
		userQueryRepository: userQueryRepository,
		tokenService:        tokenService,
	}
}

func (h *GetAllUsersQueryHandler) Handle(ctx context.Context, query *GetAllUsersQuery) (*contract.GetAllUsersResponse, error) {
	_, roles, err := h.tokenService.ValidateToken(query.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.GetAllUsersResponse](401, "Unauthorized")
		return &response, nil
	}

	isAdmin := services.IsAdminOrSuperAdmin(roles)
	hasSuperAdminRole := services.IsRole(constants.SuperAdmin, roles)

	if !isAdmin {
		response := responses.NewResponse[contract.GetAllUsersResponse](403, "Insufficient privileges")
		return &response, nil
	}

	var users []*models.UserModel

	if hasSuperAdminRole {
		users, err = h.userQueryRepository.GetUsersByRoles(constants.Admin, constants.SuperAdmin)
	} else {
		users, err = h.userQueryRepository.GetUsersByRoles(constants.Admin)
	}

	if err != nil {
		response := responses.NewResponse[contract.GetAllUsersResponse](500, "Failed to retrieve users")
		return &response, nil
	}

	userInfos := make([]models.UserSecureInfo, len(users))
	for i, user := range users {
		userInfos[i] = models.UserSecureInfo{
			ID:    user.ID,
			Roles: services.ConvertRolesToString(user.Roles),
			UserInfo: models.UserInfo{
				Name:         user.Name,
				Surname:      user.Surname,
				PhoneNumber:  user.PhoneNumber,
				EmailAddress: user.EmailAddress,
				Address:      user.Address,
				PostalCode:   user.PostalCode,
				City:         user.City,
			},
		}
	}

	return &contract.GetAllUsersResponse{
		BaseResponse: responses.NewBaseResponse(200, "Users retrieved successfully"),
		Users:        userInfos,
	}, nil
}
