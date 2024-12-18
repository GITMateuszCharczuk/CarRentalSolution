package repository_interfaces

import (
	models "blog-api/Domain/models/domestic"
	"context"
)

type BlogPostCommandRepository interface {
	CreateBlogPost(ctx context.Context, blogPost *models.BlogPostRequestModel) (*string, error)
	UpdateBlogPost(ctx context.Context, blogPost *models.BlogPostRequestModel) error
	DeleteBlogPost(ctx context.Context, id string) error
}
