package repository_interfaces

import (
	models "blog-api/Domain/models/domestic"
	"blog-api/Domain/pagination"
	"blog-api/Domain/sorting"
)

type BlogPostQueryRepository interface {
	GetBlogPostByID(id string) (*models.BlogPostResponseModel, error)
	GetBlogPosts(pagination *pagination.Pagination, sorting *sorting.Sortable,
		ids []string,
		dateTimeFrom string,
		dateTimeTo string,
		authorIds []string,
		tagsNames []string,
		visible string,
	) (*pagination.PaginatedResult[models.BlogPostResponseModel], error)
	GetBlogPostAuthorId(id string) (*string, error)
}
