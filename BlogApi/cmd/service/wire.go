// main/wire.go

// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"identity-api/API/controllers"
	"identity-api/API/server"
	validators "identity-api/API/validators"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	service_interfaces "identity-api/Domain/service_interfaces"
	config "identity-api/Infrastructure/config"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	user_mappers "identity-api/Infrastructure/databases/postgres/mappers"
	user_repository "identity-api/Infrastructure/databases/postgres/repository/user_repository"
	redis_db "identity-api/Infrastructure/databases/redis/config"
	refresh_token_repository "identity-api/Infrastructure/databases/redis/repository/refresh_token_repository"
	jwt_token_service "identity-api/Infrastructure/jwt_tocken_service"
	"identity-api/Infrastructure/password_hasher"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	Config          *config.Config
	UserQueryRepo   repository_interfaces.UserQueryRepository
	UserCommandRepo repository_interfaces.UserCommandRepository
	TokenService    service_interfaces.JWTTokenService
	PasswordHasher  service_interfaces.PasswordHasher
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		// Config
		config.WireSet,
		// Database
		postgres_db.WireSet,
		redis_db.WireSet,
		// Repository
		user_repository.WireSet,
		refresh_token_repository.WireSet,
		// Mappers
		user_mappers.WireSet,
		// Services
		jwt_token_service.WireSet,
		password_hasher.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}

func InitializeApi(userQueryRepo repository_interfaces.UserQueryRepository,
	userCommandRepo repository_interfaces.UserCommandRepository,
	tokenService service_interfaces.JWTTokenService,
	passwordHasher service_interfaces.PasswordHasher,
	config *config.Config) (*server.Server, error) {
	wire.Build(
		validators.WireSet,
		controllers.WireSet,
		server.WireSet,
	)
	return &server.Server{}, nil
}
