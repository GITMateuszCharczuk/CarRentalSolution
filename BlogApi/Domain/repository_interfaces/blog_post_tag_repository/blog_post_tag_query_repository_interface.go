package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
)

type BlogPostTagQueryRepository interface {
	GetTagByID(id string) (*models.TagModel, error)
	GetTagByName(name string) (*models.TagModel, error)
	GetTagsByBlogPostID(
		blogPostID string,
		pagination *pagination.Pagination,
	) (*pagination.PaginatedResult[models.TagModel], error)
}
