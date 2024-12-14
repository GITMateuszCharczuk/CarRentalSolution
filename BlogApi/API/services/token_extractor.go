package services

import (
	models "identity-api/Domain/models/external"

	"github.com/gin-gonic/gin"
)

func GetJwtTokenFromQuery(c *gin.Context) models.JwtToken {
	query := c.Query("token")
	return models.NewJwtToken(query)
}
