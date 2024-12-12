package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
)

type BlogPostTagCommandRepository interface {
	CreateTag(tag *models.TagModel) (*models.TagModel, error)
	UpdateTag(tag *models.TagModel) (*models.TagModel, error)
	DeleteTag(id string) error
	AddTagToBlogPost(blogPostID string, tagID string) error
	RemoveTagFromBlogPost(blogPostID string, tagID string) error
}
