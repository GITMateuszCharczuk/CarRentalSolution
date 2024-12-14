package repository_interfaces

import (
	"context"
	models "identity-api/Domain/models/domestic"
	"identity-api/Infrastructure/databases/postgres/entities"
)

type BlogPostTagCommandRepository interface {
	AddTagToBlogPost(ctx context.Context, tag *models.BlogPostTagModel, blogPostEntity entities.BlogPostEntity) (*models.BlogPostTagModel, error)
	AddTagsToBlogPost(ctx context.Context, blogPostEntity entities.BlogPostEntity, tagNames []string) error
	ModifyTagsForBlogPost(ctx context.Context, blogPostEntity entities.BlogPostEntity, newTagNames []string) error
	CleanupUnusedTags(ctx context.Context) error
}
