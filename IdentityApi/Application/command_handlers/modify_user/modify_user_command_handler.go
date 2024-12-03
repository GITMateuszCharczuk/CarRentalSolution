package commands

import (
	"context"
	contract "identity-api/Application.contract/modify_user"
	"identity-api/Application/services"
	models "identity-api/Domain/models/user"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
	"time"
)

type ModifyUserCommandHandler struct {
	userQueryRepository   repository_interfaces.UserQueryRepository
	userCommandRepository repository_interfaces.UserCommandRepository
	tokenService          service_interfaces.JWTTokenService
}

func NewModifyUserCommandHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	userCommandRepository repository_interfaces.UserCommandRepository,
	tokenService service_interfaces.JWTTokenService,
) *ModifyUserCommandHandler {
	return &ModifyUserCommandHandler{
		userQueryRepository:   userQueryRepository,
		userCommandRepository: userCommandRepository,
		tokenService:          tokenService,
	}
}

func (h *ModifyUserCommandHandler) Handle(ctx context.Context, command *ModifyUserCommand) (*contract.ModifyUserResponse, error) {
	requesterID, requesterRoles, err := h.tokenService.ValidateToken(command.JwtToken)
	var existingUser *models.UserModel
	if err != nil {
		return &contract.ModifyUserResponse{
			BaseResponse: responses.NewBaseResponse(401, "Unauthorized"),
		}, nil
	}

	isAdmin := services.IsAdminOrSuperAdmin(requesterRoles)

	if command.UserID != "" && isAdmin {
		existingUser, err = h.userQueryRepository.GetUserByID(command.UserID)
		if err != nil || existingUser == nil {
			return &contract.ModifyUserResponse{
				BaseResponse: responses.NewBaseResponse(404, "User not found"),
			}, nil
		}
	} else {
		existingUser, err = h.userQueryRepository.GetUserByID(requesterID)
		if err != nil || existingUser == nil {
			return &contract.ModifyUserResponse{
				BaseResponse: responses.NewBaseResponse(404, "User not found"),
			}, nil
		}
	}

	existingUser.Name = command.Name
	existingUser.Surname = command.Surname
	existingUser.PhoneNumber = command.PhoneNumber
	existingUser.Address = command.Address
	existingUser.PostalCode = command.PostalCode
	existingUser.City = command.City
	existingUser.UpdatedAt = time.Now()

	if isAdmin && len(command.Roles) > 0 {
		existingUser.Roles = services.ConvertRolesToJWTRole(command.Roles)
	}

	if err := h.userCommandRepository.UpdateUser(existingUser); err != nil {
		return &contract.ModifyUserResponse{
			BaseResponse: responses.NewBaseResponse(500, "Failed to update user"),
		}, nil
	}

	return &contract.ModifyUserResponse{
		BaseResponse: responses.NewBaseResponse(200, "User updated successfully"),
	}, nil
}
