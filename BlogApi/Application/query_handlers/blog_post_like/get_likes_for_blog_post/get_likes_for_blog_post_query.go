package queries

type GetLikesForBlogPostQuery struct {
	BlogPostId string `json:"blogPostId" validate:"required"`
}
