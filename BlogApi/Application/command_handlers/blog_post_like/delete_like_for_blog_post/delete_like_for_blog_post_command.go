package commands

import (
	models "identity-api/Domain/models/external"
)

type DeleteLikeForBlogPostCommand struct {
	BlogPostId string `json:"blogPostId" validate:"required"`
	models.JwtToken
}
