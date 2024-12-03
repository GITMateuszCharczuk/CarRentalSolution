package models

type JwtToken struct {
	Token string `json:"token" binding:"required" example:"jwt.token.here" swaggertype:"string"`
}

func NewJwtToken(token string) JwtToken {
	return JwtToken{
		Token: token,
	}
}
