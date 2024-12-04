package commands

import (
	delete "identity-api/Application/command_handlers/delete_user"
	login "identity-api/Application/command_handlers/login"
	modify "identity-api/Application/command_handlers/modify_user"
	refresh_token "identity-api/Application/command_handlers/refresh_token"
	register "identity-api/Application/command_handlers/register"
	validate_token "identity-api/Application/command_handlers/validate_token"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	service_interfaces "identity-api/Domain/service_interfaces"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerRegisterCommandHandler(
	hasher service_interfaces.PasswordHasher,
	userQueryRepository repository_interfaces.UserQueryRepository,
	userCommandRepository repository_interfaces.UserCommandRepository,
) {
	handler := register.NewRegisterCommandHandler(hasher, userQueryRepository, userCommandRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerModifyUserCommandHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	userCommandRepository repository_interfaces.UserCommandRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	handler := modify.NewModifyUserCommandHandler(userQueryRepository, userCommandRepository, tokenService)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteUserCommandHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	userCommandRepository repository_interfaces.UserCommandRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	handler := delete.NewDeleteUserCommandHandler(userQueryRepository, userCommandRepository, tokenService)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerValidateTokenCommandHandler(
	tokenService service_interfaces.JWTTokenService,
) {
	handler := validate_token.NewValidateTokenCommandHandler(tokenService)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerLoginCommandHandler(
	hasher service_interfaces.PasswordHasher,
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	handler := login.NewLoginCommandHandler(hasher, tokenService, userQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerRefreshTokenCommandHandler(
	tokenService service_interfaces.JWTTokenService,
) {
	handler := refresh_token.NewRefreshTokenCommandHandler(tokenService)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterCommandHandlers(
	hasher service_interfaces.PasswordHasher,
	userQueryRepository repository_interfaces.UserQueryRepository,
	userCommandRepository repository_interfaces.UserCommandRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	registerRegisterCommandHandler(hasher, userQueryRepository, userCommandRepository)
	registerModifyUserCommandHandler(userQueryRepository, userCommandRepository, tokenService)
	registerDeleteUserCommandHandler(userQueryRepository, userCommandRepository, tokenService)
	registerValidateTokenCommandHandler(tokenService)
	registerLoginCommandHandler(hasher, userQueryRepository, tokenService)
	registerRefreshTokenCommandHandler(tokenService)
}
