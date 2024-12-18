package commands

import (
	models "blog-api/Domain/models/external"
)

type DeleteLikeForBlogPostCommand struct {
	BlogPostId string `json:"blogPostId" validate:"required"`
	models.JwtToken
}
