package commands

import (
	"context"
	contract "identity-api/Application.contract/delete_user"
	"identity-api/Domain/constants"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
)

type DeleteUserCommandHandler struct {
	userQueryRepository   repository_interfaces.UserQueryRepository
	userCommandRepository repository_interfaces.UserCommandRepository
	tokenService          service_interfaces.JWTTokenService
}

func NewDeleteUserCommandHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	userCommandRepository repository_interfaces.UserCommandRepository,
	tokenService service_interfaces.JWTTokenService,
) *DeleteUserCommandHandler {
	return &DeleteUserCommandHandler{
		userQueryRepository:   userQueryRepository,
		userCommandRepository: userCommandRepository,
		tokenService:          tokenService,
	}
}

func (h *DeleteUserCommandHandler) Handle(ctx context.Context, command *DeleteUserCommand) (*contract.DeleteUserResponse, error) {
	_, requesterRoles, err := h.tokenService.ValidateToken(command.JwtToken)
	if err != nil {
		return &contract.DeleteUserResponse{
			BaseResponse: responses.NewBaseResponse(401, "Unauthorized"),
		}, nil
	}

	isSuperAdmin := false
	isAdmin := false
	for _, role := range requesterRoles {
		if role == constants.SuperAdmin {
			isSuperAdmin = true
			break
		}
		if role == constants.Admin {
			isAdmin = true
		}
	}

	if !isAdmin && !isSuperAdmin {
		return &contract.DeleteUserResponse{
			BaseResponse: responses.NewBaseResponse(403, "Insufficient privileges"),
		}, nil
	}

	userToDelete, err := h.userQueryRepository.GetUserByID(command.ID)
	if err != nil || userToDelete == nil {
		return &contract.DeleteUserResponse{
			BaseResponse: responses.NewBaseResponse(404, "User not found"),
		}, nil
	}

	for _, role := range userToDelete.Roles {
		if role == constants.SuperAdmin && !isSuperAdmin {
			return &contract.DeleteUserResponse{
				BaseResponse: responses.NewBaseResponse(403, "Cannot delete SuperAdmin user"),
			}, nil
		}
	}

	if err := h.userCommandRepository.DeleteUser(command.ID); err != nil {
		return &contract.DeleteUserResponse{
			BaseResponse: responses.NewBaseResponse(500, "Failed to delete user"),
		}, nil
	}

	return &contract.DeleteUserResponse{
		BaseResponse: responses.NewBaseResponse(200, "User deleted successfully"),
	}, nil
}
