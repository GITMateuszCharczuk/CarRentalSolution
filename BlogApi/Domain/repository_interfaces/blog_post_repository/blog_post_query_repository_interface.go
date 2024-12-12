package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	"identity-api/Domain/sorting"
)

type BlogPostQueryRepository interface {
	GetBlogPostByID(id string) (*models.BlogPostModel, error)
	GetBlogPostByUrlHandle(urlHandle string) (*models.BlogPostModel, error)
	GetBlogPosts(pagination *pagination.Pagination, sorting *sorting.Sortable) (*pagination.PaginatedResult[models.BlogPostModel], error)
	GetBlogPostsByAuthorID(authorID string, pagination *pagination.Pagination, sorting *sorting.Sortable) (*pagination.PaginatedResult[models.BlogPostModel], error)
	GetVisibleBlogPosts(pagination *pagination.Pagination, sorting *sorting.Sortable) (*pagination.PaginatedResult[models.BlogPostModel], error)
}
