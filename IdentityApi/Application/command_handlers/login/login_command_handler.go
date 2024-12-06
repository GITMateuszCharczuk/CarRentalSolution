package commands

import (
	"context"
	contract "identity-api/Application.contract/login"
	services "identity-api/Application/services"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
	"log"
)

type LoginCommandHandler struct {
	hasher              service_interfaces.PasswordHasher
	tokenService        service_interfaces.JWTTokenService
	userQueryRepository repository_interfaces.UserQueryRepository
}

func NewLoginCommandHandler(hasher service_interfaces.PasswordHasher, tokenService service_interfaces.JWTTokenService, userQueryRepository repository_interfaces.UserQueryRepository) *LoginCommandHandler {
	return &LoginCommandHandler{hasher: hasher, tokenService: tokenService, userQueryRepository: userQueryRepository}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, command *LoginCommand) (*contract.LoginResponse, error) {
	notAuthorized := responses.NewResponse[contract.LoginResponse](401, "Invalid email or password")
	user, err := h.userQueryRepository.GetUserByEmail(command.Email)
	if err != nil || user == nil {
		log.Println(err)
		return &notAuthorized, nil
	}
	if valid, err := h.hasher.VerifyPassword(user.Password, command.Password); err != nil || !valid {
		return &notAuthorized, nil
	}

	token, refreshToken, err := h.tokenService.GenerateTokens(user.ID, user.Roles)
	if err != nil {
		return &notAuthorized, nil
	}

	return &contract.LoginResponse{
		BaseResponse:    responses.NewBaseResponse(200, "Login successful"),
		JwtToken:        token,
		JwtRefreshToken: refreshToken,
		Roles:           services.ConvertRolesToString(user.Roles),
	}, nil
}
