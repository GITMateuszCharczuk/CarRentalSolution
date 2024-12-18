package contract

import models "blog-api/Domain/models/external"

type DeleteLikeForBlogPostRequest struct {
	BlogPostId      string `json:"blogPostId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	models.JwtToken `json:",inline"`
}
