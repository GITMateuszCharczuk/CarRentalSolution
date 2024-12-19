package contract

type GetTagsRequest struct {
	SortQuery  []string `json:"sort_query" validate:"validBlogPostTagSortable"`
	CarOfferId string   `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"omitempty,len=36"`
}
