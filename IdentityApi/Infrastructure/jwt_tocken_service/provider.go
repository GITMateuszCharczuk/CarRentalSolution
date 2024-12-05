package jwt_token_service

import (
	repository_interfaces "identity-api/Domain/repository_interfaces/refresh_token_repository"
	service_interfaces "identity-api/Domain/service_interfaces"
	"identity-api/Infrastructure/config"
	"time"

	"github.com/google/wire"
)

func ProvideJWTTokenService(cfg *config.Config, commandRepo repository_interfaces.RefreshTokenCommandRepository, queryRepo repository_interfaces.RefreshTokenQueryRepository) service_interfaces.JWTTokenService {
	accessTokenTTL := time.Duration(cfg.AccessTokenTTL) * time.Minute
	refreshTokenTTL := time.Duration(cfg.RefreshTokenTTL) * time.Minute
	return NewJWTTokenService(accessTokenTTL, refreshTokenTTL, cfg.SecretKey, commandRepo, queryRepo)
}

var WireSet = wire.NewSet(ProvideJWTTokenService)
