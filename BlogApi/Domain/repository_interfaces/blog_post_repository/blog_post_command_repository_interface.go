package repository_interfaces

import (
	"context"
	models "identity-api/Domain/models/domestic"
)

type BlogPostCommandRepository interface {
	CreateBlogPost(ctx context.Context, blogPost *models.BlogPostRequestModel) (*string, error)
	UpdateBlogPost(ctx context.Context, blogPost *models.BlogPostRequestModel) error
	DeleteBlogPost(ctx context.Context, id string) error
}
