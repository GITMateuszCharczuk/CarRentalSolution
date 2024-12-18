package commands

import (
	models "blog-api/Domain/models/external"
)

type CreateLikeForBlogPostCommand struct {
	BlogPostId string `json:"blogPostId" validate:"required"`
	models.JwtToken
}
