package repository_interfaces

import (
	models "blog-api/Domain/models/domestic"
	"blog-api/Infrastructure/databases/postgres/entities"
	"context"
)

type BlogPostTagCommandRepository interface {
	AddTagToBlogPost(ctx context.Context, tag *models.BlogPostTagModel, blogPostEntity entities.BlogPostEntity) (*models.BlogPostTagModel, error)
	AddTagsToBlogPost(ctx context.Context, blogPostId string, tagNames []string) error
	ModifyTagsForBlogPost(ctx context.Context, blogPostId string, newTagNames []string) error
	CleanupUnusedTags(ctx context.Context) error
}
