package service_interfaces

import (
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/token"
)

type JWTTokenService interface {
	GenerateTokens(userID string, roles []constants.JWTRole) (models.JwtToken, models.JwtRefreshToken, error)

	ValidateToken(token models.JwtToken) (string, []constants.JWTRole, error)

	RefreshToken(refreshToken models.JwtRefreshToken) (models.JwtToken, error)

	RevokeToken(token models.JwtRefreshToken) error
}
