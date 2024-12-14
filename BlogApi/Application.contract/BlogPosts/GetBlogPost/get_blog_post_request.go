package contract

type GetBlogPostRequest struct {
	BlogPostId string `json:"blogPostId" validate:"required"`
}
