package contract

type TagSortColumn string
type SortOrder string

const (
	SortByName TagSortColumn = "name"
	SortById   TagSortColumn = "id"

	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type GetTagsRequest struct {
	OrderBy        *TagSortColumn `json:"orderBy" example:"name" swaggertype:"string" enums:"name,id"`
	OrderDirection *SortOrder     `json:"orderDirection" example:"asc" swaggertype:"string" enums:"asc,desc"`
	BlogPostId     string         `json:"blogPostId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}
