package repository_interfaces

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	"identity-api/Domain/sorting"
)

type BlogPostQueryRepository interface {
	GetBlogPostByID(id string) (*models.BlogPostResponseModel, error)
	GetBlogPosts(pagination *pagination.Pagination, sorting *sorting.Sortable,
		ids []string,
		dateTimeFrom string,
		dateTimeTo string,
		authorIds []string,
		tagsNames []string,
		visible bool,
	) (*pagination.PaginatedResult[models.BlogPostResponseModel], error)
	GetBlogPostAuthorId(id string) (*string, error)
}
