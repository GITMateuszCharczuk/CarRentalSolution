package service_interfaces

import (
	"identity-api/Domain/constants"
	"identity-api/Domain/models"
	"time"
)

type JWTTokenService interface {
	GenerateTokens(userID string, roles []constants.JWTRole, expiration time.Duration) (models.JwtToken, models.JwtRefreshToken, error)

	ValidateToken(token models.JwtToken) (string, []constants.JWTRole, error)

	RefreshToken(refreshToken models.JwtRefreshToken) (models.JwtToken, error)

	RevokeToken(token models.JwtRefreshToken) error
}
