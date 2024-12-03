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
		return &contract.GetAllUsersResponse{
			BaseResponse: responses.NewBaseResponse(401, "Unauthorized"),
		}, nil
	}

	hasAdminRole := false
	for _, role := range roles {
		if role == constants.Admin || role == constants.SuperAdmin {
			hasAdminRole = true
			break
		}
	}

	if !hasAdminRole {
		return &contract.GetAllUsersResponse{
			BaseResponse: responses.NewBaseResponse(403, "Insufficient privileges"),
		}, nil
	}

	users, err := h.userQueryRepository.GetAllUsers()
	if err != nil {
		return &contract.GetAllUsersResponse{
			BaseResponse: responses.NewBaseResponse(500, "Failed to retrieve users"),
		}, nil
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
