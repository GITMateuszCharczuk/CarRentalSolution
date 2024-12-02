package repository_interfaces

type RefreshTokenRepository interface {
	SaveRefreshToken(userID string, refreshToken string, ttl int) error

	GetRefreshToken(refreshToken string) (string, error)

	RevokeRefreshToken(refreshToken string) error
}
