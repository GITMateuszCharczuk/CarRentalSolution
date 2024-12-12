package repository

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	"identity-api/Domain/sorting"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
	"identity-api/Infrastructure/databases/postgres/repository/base/helpers"
)

type BlogPostCommentQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel]
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel]
}

func NewBlogPostCommentQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
) repository_interfaces.BlogPostCommentQueryRepository {
	return &BlogPostCommentQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel](postgresDatabase.DB, mapper),
		mapper:          mapper,
	}
}

func (r *BlogPostCommentQueryRepositoryImpl) GetCommentByID(id string) (*models.BlogPostCommentModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.BlogPostCommentEntity]("id", id)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *BlogPostCommentQueryRepositoryImpl) GetComments(
	blogPostIDs []string,
	userIDs []string,
	dateTimeFrom string,
	dateTimeTo string,
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
) (*pagination.PaginatedResult[models.BlogPostCommentModel], error) {
	queryRecords := []helpers.QueryRecord[entities.BlogPostCommentEntity]{
		helpers.NewQueryRecord[entities.BlogPostCommentEntity]("blogPostID", blogPostIDs),
		helpers.NewQueryRecord[entities.BlogPostCommentEntity]("userID", userIDs),
	}

	query := r.ConstructBaseQuery()
	query = r.EnrichQuery(query, queryRecords...)

	if dateTimeFrom != "" && dateTimeTo != "" {
		query = query.Where("created_at BETWEEN ? AND ?", dateTimeFrom, dateTimeTo)
	} else if dateTimeFrom != "" {
		query = query.Where("created_at >= ?", dateTimeFrom)
	} else if dateTimeTo != "" {
		query = query.Where("created_at <= ?", dateTimeTo)
	}

	return r.ExecutePaginatedQuery(query, pagination, sorting)
}
