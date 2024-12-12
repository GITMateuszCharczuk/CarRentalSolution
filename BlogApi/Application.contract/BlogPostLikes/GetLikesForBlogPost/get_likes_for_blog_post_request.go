package contract

type GetLikesForBlogPostRequest struct {
	BlogPostId string `json:"blogPostId" validate:"required"`
}
