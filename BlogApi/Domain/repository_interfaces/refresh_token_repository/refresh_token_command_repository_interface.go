package repository_interfaces

type RefreshTokenCommandRepository interface {
	SaveRefreshToken(userID string, refreshToken string, ttl int) error
	RevokeRefreshToken(refreshToken string) error
}
