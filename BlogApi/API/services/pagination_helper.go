package services

import (
	"identity-api/Domain/pagination"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractPagination(c *gin.Context) pagination.Pagination {
	var pagination pagination.Pagination

	if pageSize := c.Query("page_size"); pageSize != "" {
		if ps, err := strconv.Atoi(pageSize); err == nil {
			pagination.PageSize = ps
			pagination.Enabled = true
		}
	}

	if currentPage := c.Query("current_page"); currentPage != "" {
		if cp, err := strconv.Atoi(currentPage); err == nil {
			pagination.CurrentPage = cp
			pagination.Enabled = true
		}
	}

	return pagination
}
