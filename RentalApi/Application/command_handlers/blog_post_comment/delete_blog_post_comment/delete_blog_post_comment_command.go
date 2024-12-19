package commands

import (
	models "rental-api/Domain/models/external"
)

type DeleteBlogPostCommentCommand struct {
	BlogPostCommentId string `json:"blogPostCommentId" validate:"required"`
	models.JwtToken
}
