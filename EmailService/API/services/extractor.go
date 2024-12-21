package services

import (
	models "email-service/Domain/models/external"
	pagination "email-service/Domain/requests"
	"strconv"

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

func ExtractPagination(c *gin.Context) pagination.Pagination {
	var pagination pagination.Pagination

	if pageSize := c.Query("page_size"); pageSize != "" {
		if ps, err := strconv.Atoi(pageSize); err == nil {
			pagination.PageSize = ps
		}
	}

	if page := c.Query("page"); page != "" {
		if cp, err := strconv.Atoi(page); err == nil {
			pagination.Page = cp
		}
	}

	return pagination
}
