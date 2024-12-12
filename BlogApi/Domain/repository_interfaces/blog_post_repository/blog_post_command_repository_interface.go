package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
)

type BlogPostCommandRepository interface {
	CreateBlogPost(blogPost *models.BlogPostModel) (*models.BlogPostModel, error)
	UpdateBlogPost(blogPost *models.BlogPostModel) (*models.BlogPostModel, error)
	DeleteBlogPost(id string) error
}
