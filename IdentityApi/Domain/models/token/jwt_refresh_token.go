package models

type JwtRefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"jwt.refresh.token.here" swaggerignore:"true"`
}

func NewRefreshToken(refreshToken string) JwtRefreshToken {

	return JwtRefreshToken{
		RefreshToken: refreshToken,
	}
}
