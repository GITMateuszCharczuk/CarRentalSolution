package services

import (
	models "file-storage/Domain/models/external"

	"github.com/gin-gonic/gin"
)

func GetJwtTokenFromQuery(c *gin.Context) models.JwtToken {
	query := c.Query("token")
	return models.NewJwtToken(query)
}

func ExtractFromPath(c *gin.Context, key string) string {
	value := c.Param(key)

	if value == "undefined" {
		return ""
	}

	return value
}
