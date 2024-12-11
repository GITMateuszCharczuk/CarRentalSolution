package services

import (
	token "identity-api/Domain/models/token"

	"github.com/gin-gonic/gin"
)

func GetJwtTokenFromQuery(c *gin.Context) token.JwtToken {
	query := c.Query("token")
	return token.NewJwtToken(query)
}

func GetJwtRefreshTokenFromQuery(c *gin.Context) token.JwtRefreshToken {
	query := c.Query("token")
	return token.NewRefreshToken(query)
}
