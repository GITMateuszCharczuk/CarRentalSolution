package contract

type GetBlogPostRequest struct {
	BlogPostId string `json:"blogPostId" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}
