package repository_interfaces

import (
	"context"
	models "identity-api/Domain/models/domestic"
)

type BlogPostCommentCommandRepository interface {
	AddComment(ctx context.Context, comment *models.BlogPostCommentModel) (string, error)
	UpdateComment(ctx context.Context, comment *models.BlogPostCommentModel) error
	RemoveComment(ctx context.Context, commentID string) error
}
