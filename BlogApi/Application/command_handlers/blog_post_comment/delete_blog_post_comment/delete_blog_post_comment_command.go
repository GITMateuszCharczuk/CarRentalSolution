package commands

import (
	models "identity-api/Domain/models/external"
)

type DeleteBlogPostCommentCommand struct {
	BlogPostCommentId string `json:"blogPostCommentId" validate:"required"`
	models.JwtToken
}
