package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
)

type BlogPostCommentCommandRepository interface {
	AddComment(comment *models.BlogPostCommentModel) (string, error)
	UpdateComment(comment *models.BlogPostCommentModel) error
	RemoveComment(commentID string) error
}
