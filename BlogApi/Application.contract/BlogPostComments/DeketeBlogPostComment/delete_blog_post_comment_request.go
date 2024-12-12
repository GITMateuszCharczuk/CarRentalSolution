package contract

import models "identity-api/Domain/models/external"

type DeleteBlogPostCommentRequest struct {
	BlogPostCommentId string `json:"blogPostCommentId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	models.JwtToken
}
