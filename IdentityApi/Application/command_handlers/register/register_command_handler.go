package commands

import (
	"context"
	contract "identity-api/Application.contract/register"
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/user"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
	"time"

	"github.com/google/uuid"
)

type RegisterCommandHandler struct {
	hasher                service_interfaces.PasswordHasher
	userQueryRepository   repository_interfaces.UserQueryRepository
	userCommandRepository repository_interfaces.UserCommandRepository
}

func NewRegisterCommandHandler(
	hasher service_interfaces.PasswordHasher,
	userQueryRepository repository_interfaces.UserQueryRepository,
	userCommandRepository repository_interfaces.UserCommandRepository,
) *RegisterCommandHandler {
	return &RegisterCommandHandler{
		hasher:                hasher,
		userQueryRepository:   userQueryRepository,
		userCommandRepository: userCommandRepository,
	}
}

func (h *RegisterCommandHandler) Handle(ctx context.Context, command *contract.RegisterUserRequest) (*contract.RegisterUserResponse, error) {
	existingUser, _ := h.userQueryRepository.GetUserByEmail(command.EmailAddress)
	if existingUser != nil {
		return &contract.RegisterUserResponse{
			BaseResponse: responses.NewBaseResponse(400, "Email already exists"),
		}, nil
	}

	hashedPassword, err := h.hasher.HashPassword(command.Password)
	if err != nil {
		return &contract.RegisterUserResponse{
			BaseResponse: responses.NewBaseResponse(500, "Error processing registration"),
		}, nil
	}

	user := &models.UserModel{
		ID:           uuid.New().String(),
		Roles:        []constants.JWTRole{constants.User},
		Name:         command.Name,
		Surname:      command.Surname,
		PhoneNumber:  command.PhoneNumber,
		EmailAddress: command.EmailAddress,
		Password:     hashedPassword,
		Address:      command.Address,
		PostalCode:   command.PostalCode,
		City:         command.City,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := h.userCommandRepository.CreateUser(user); err != nil {
		return &contract.RegisterUserResponse{
			BaseResponse: responses.NewBaseResponse(500, "Failed to create user"),
		}, nil
	}

	return &contract.RegisterUserResponse{
		BaseResponse: responses.NewBaseResponse(201, "User registered successfully"),
	}, nil
}
