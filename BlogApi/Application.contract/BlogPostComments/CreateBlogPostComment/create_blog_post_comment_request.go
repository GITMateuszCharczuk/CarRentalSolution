package contract

import models "blog-api/Domain/models/external"

type CreateBlogPostCommentRequest struct {
	Description string `json:"description" example:"This is a great blog post!" swaggertype:"string" validate:"required"`
	BlogPostId  string `json:"-" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true" validate:"required"`
	models.JwtToken
}
