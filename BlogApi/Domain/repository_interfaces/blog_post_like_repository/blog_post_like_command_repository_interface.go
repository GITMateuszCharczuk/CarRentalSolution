package repository_interfaces

import "context"

type BlogPostLikeCommandRepository interface {
	AddLike(ctx context.Context, blogPostID string, userID string) error
	RemoveLike(ctx context.Context, blogPostID string, userID string) error
}
