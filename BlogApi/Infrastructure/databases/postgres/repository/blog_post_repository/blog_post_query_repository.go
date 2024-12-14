package repository

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	"identity-api/Domain/sorting"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
	"identity-api/Infrastructure/databases/postgres/repository/base/helpers"
)

type BlogPostQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostEntity, string, models.BlogPostResponseModel]
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostResponseModel]
}

func NewBlogPostQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostResponseModel],
) repository_interfaces.BlogPostQueryRepository {
	return &BlogPostQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostEntity, string, models.BlogPostResponseModel](postgresDatabase.DB, mapper),
		mapper:          mapper,
	}
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPostByID(id string) (*models.BlogPostResponseModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.BlogPostEntity]("id", id)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPosts(
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
	ids []string,
	dateTimeFrom string,
	dateTimeTo string,
	authorIds []string,
	tagNames []string,
	visible bool,
) (*pagination.PaginatedResult[models.BlogPostResponseModel], error) {
	query := r.ConstructBaseQuery()
	if len(tagNames) > 0 {
		query = query.Joins("Tags").Where("Tags.name IN ?", tagNames)
	}

	queryRecords := []helpers.QueryRecord[entities.BlogPostEntity]{
		helpers.NewQueryRecord[entities.BlogPostEntity]("id", ids),
		helpers.NewQueryRecord[entities.BlogPostEntity]("authorID", authorIds),
		helpers.NewQueryRecord[entities.BlogPostEntity]("visible", visible),
	}
	query = r.ApplyWhereConditions(query, queryRecords...)

	if dateTimeFrom != "" && dateTimeTo != "" {
		query = query.Where("blog_post_entities.created_at BETWEEN ? AND ?", dateTimeFrom, dateTimeTo)
	} else if dateTimeFrom != "" {
		query = query.Where("blog_post_entities.created_at >= ?", dateTimeFrom)
	} else if dateTimeTo != "" {
		query = query.Where("blog_post_entities.created_at <= ?", dateTimeTo)
	}

	return r.ExecutePaginatedQuery(query, pagination, sorting)
}
