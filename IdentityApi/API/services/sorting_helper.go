package services

import (
	"identity-api/Domain/sorting"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractSorting(sortQuery []string) sorting.Sortable {
	var sortable sorting.Sortable

	if sortFields := sortQuery; len(sortFields) > 0 {
		sortable.Enabled = false
		sortable.SortFields = make([]sorting.SortField, 0)

		for _, sf := range sortFields {
			parts := strings.Split(sf, ":")
			if len(parts) == 2 {
				var direction sorting.SortDirection
				if strings.ToLower(parts[1]) == "desc" {
					direction = sorting.DESC
					sortable.Enabled = true
				}
				if strings.ToLower(parts[1]) == "asc" {
					direction = sorting.ASC
					sortable.Enabled = true
				}
				sortable.SortFields = append(sortable.SortFields, sorting.SortField{
					Field:     parts[0],
					Direction: direction,
				})
			}
		}
	}
	return sortable
}

func ExtractSortQuery(c *gin.Context) []string {
	return c.QueryArray("sort_fields")
}
