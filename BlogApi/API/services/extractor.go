package services

import (
	models "identity-api/Domain/models/external"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetJwtTokenFromQuery(c *gin.Context) models.JwtToken {
	query := c.Query("token")
	return models.NewJwtToken(query)
}

func ExtractQueryArray(c *gin.Context, key string) []string {
	values := c.QueryArray(key)

	if len(values) == 1 && strings.Contains(values[0], ",") {
		values = strings.Split(values[0], ",")
	}

	for i, value := range values {
		values[i] = strings.TrimSpace(value)
	}

	return values
}
