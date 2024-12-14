package commands

import (
	models "identity-api/Domain/models/external"
)

type CreateBlogPostCommentCommand struct {
	Description string `json:"description" validate:"required"`
	BlogPostId  string `json:"blogPostId" validate:"required"`
	models.JwtToken
}
