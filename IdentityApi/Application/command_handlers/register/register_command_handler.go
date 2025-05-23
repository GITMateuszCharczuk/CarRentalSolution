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

func (h *RegisterCommandHandler) Handle(ctx context.Context, command *RegisterUserCommand) (*contract.RegisterUserResponse, error) {
	existingUser, _ := h.userQueryRepository.GetUserByEmail(command.EmailAddress)
	if existingUser != nil {
		response := responses.NewResponse[contract.RegisterUserResponse](400, "Email already exists")
		return &response, nil
	}

	hashedPassword, err := h.hasher.HashPassword(command.Password)
	if err != nil {
		response := responses.NewResponse[contract.RegisterUserResponse](500, "Error processing registration")
		return &response, nil
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
		response := responses.NewResponse[contract.RegisterUserResponse](500, "Failed to create user")
		return &response, nil
	}

	response := responses.NewResponse[contract.RegisterUserResponse](201, "User registered successfully")
	return &response, nil
}
