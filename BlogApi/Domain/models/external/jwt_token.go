package models

type JwtToken struct {
	Token string `json:"token" binding:"required" example:"your.jwt.token.here" swaggerignore:"true"`
}

func NewJwtToken(token string) JwtToken {
	return JwtToken{
		Token: token,
	}
}
