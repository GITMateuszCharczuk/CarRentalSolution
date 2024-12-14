package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/sorting"
)

type BlogPostTagQueryRepository interface {
	GetTagByID(id string) (*models.BlogPostTagModel, error)
	GetTagByName(name string) (*models.BlogPostTagModel, error)
	GetTagsByBlogPostID(
		blogPostID string,
		sorting sorting.Sortable,
	) (*[]models.BlogPostTagModel, error)
}
