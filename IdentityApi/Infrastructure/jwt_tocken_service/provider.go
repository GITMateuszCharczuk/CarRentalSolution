package jwt_token_service

import (
	"identity-api/Domain/repository_interfaces"
	"identity-api/Infrastructure/config"
)

import "time"

func ProvideJWTTokenService(cfg *config.Config, refreshTokenRepo repository_interfaces.RefreshTokenRepository) *JWTTokenServiceImpl {
	accessTokenTTL := time.Duration(cfg.AccessTokenTTL) * time.Minute
	refreshTokenTTL := time.Duration(cfg.RefreshTokenTTL) * time.Minute
	return NewJWTTokenService(accessTokenTTL, refreshTokenTTL, cfg.SecretKey, refreshTokenRepo)
}
