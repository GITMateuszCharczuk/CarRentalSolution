package repository_interfaces

import (
	models "blog-api/Domain/models/domestic"
	"context"
)

type BlogPostCommentCommandRepository interface {
	AddComment(ctx context.Context, comment *models.BlogPostCommentModel) (string, error)
	UpdateComment(ctx context.Context, comment *models.BlogPostCommentModel) error
	RemoveComment(ctx context.Context, commentID string) error
}
