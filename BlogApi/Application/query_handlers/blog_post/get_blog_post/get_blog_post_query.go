package queries

type GetBlogPostQuery struct {
	ID string `json:"id" validate:"required"`
}
