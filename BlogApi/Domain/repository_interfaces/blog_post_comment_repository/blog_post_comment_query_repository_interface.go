package repository_interfaces

import (
	models "blog-api/Domain/models/domestic"
	"blog-api/Domain/pagination"
	"blog-api/Domain/sorting"
)

type BlogPostCommentQueryRepository interface {
	GetCommentByID(id string) (*models.BlogPostCommentModel, error)
	GetCommentAuthorId(id string) (*string, error)
	GetComments(
		blogPostIDs []string,
		userIDs []string,
		dateTimeFrom string,
		dateTimeTo string,
		pagination *pagination.Pagination,
		sorting *sorting.Sortable,
	) (*pagination.PaginatedResult[models.BlogPostCommentModel], error)
	GetCommentsCount(blogPostId string) (int, error)
}
