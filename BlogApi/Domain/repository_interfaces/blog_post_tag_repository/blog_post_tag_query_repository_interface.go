package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	"identity-api/Domain/sorting"
)

type BlogPostTagQueryRepository interface {
	GetTagByID(id string) (*models.TagModel, error)
	GetTagByName(name string) (*models.TagModel, error)
	GetAllTags(pagination *pagination.Pagination, sorting *sorting.Sortable) (*pagination.PaginatedResult[models.TagModel], error)
	GetTagsByBlogPostID(blogPostID string) ([]models.TagModel, error)
	GetBlogPostsByTag(tagName string, pagination *pagination.Pagination, sorting *sorting.Sortable) (*pagination.PaginatedResult[models.BlogPostModel], error)
}
