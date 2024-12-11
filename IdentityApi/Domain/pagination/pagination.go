package pagination

type Pagination struct {
	PageSize    int  `json:"page_size" validate:"required_if=Enabled true,min=1,max=100" example:"10" swaggerignore:"false"`
	CurrentPage int  `json:"current_page" validate:"required_if=Enabled true,min=1" example:"1" swaggerignore:"false"`
	Enabled     bool `json:"-" swaggerignore:"true"`
}

type PaginatedResult[T any] struct {
	Items       []T   `json:"items"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
}
