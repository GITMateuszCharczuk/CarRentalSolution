package contract

type CreateBlogPostCommentRequest struct {
	Description string `json:"description" binding:"required" example:"This is a great blog post!" swaggertype:"string" validate:"required"`
	BlogPostId  string `json:"blogPostId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	UserId      string `json:"userId" binding:"required" example:"456e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"
}
