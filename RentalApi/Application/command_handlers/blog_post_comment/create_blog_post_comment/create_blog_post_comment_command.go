package commands

import (
	models "rental-api/Domain/models/external"
)

type CreateBlogPostCommentCommand struct {
	Description string `json:"description" validate:"required"`
	BlogPostId  string `json:"blogPostId" validate:"required"`
	models.JwtToken
}
