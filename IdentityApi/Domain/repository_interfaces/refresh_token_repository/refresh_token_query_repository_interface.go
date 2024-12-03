package repository_interfaces

type RefreshTokenQueryRepository interface {
	GetRefreshToken(refreshToken string) (string, error)
}
