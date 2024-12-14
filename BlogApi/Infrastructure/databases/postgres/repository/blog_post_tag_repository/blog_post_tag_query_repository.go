package repository

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_tag_repository"
	"identity-api/Domain/sorting"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
	"identity-api/Infrastructure/databases/postgres/repository/base/helpers"
)

type BlogPostTagQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostTagEntity, string, models.TagModel]
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.TagModel]
}

func NewBlogPostTagQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.TagModel],
) repository_interfaces.BlogPostTagQueryRepository {
	return &BlogPostTagQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostTagEntity, string, models.TagModel](postgresDatabase.DB, mapper),
		mapper:          mapper,
	}
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagByID(id string) (*models.TagModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.BlogPostTagEntity]("id", id)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagByName(name string) (*models.TagModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.BlogPostTagEntity]("name", name)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagsByBlogPostID(
	blogPostID string,
	pagination *pagination.Pagination,
) (*pagination.PaginatedResult[models.TagModel], error) {
	var sorting = sorting.Sortable{}
	query := r.DbContext.Joins("JOIN blog_post_tags ON blog_post_tags.tag_id = blog_post_tag_entities.id").
		Where("blog_post_tags.blog_post_id = ?", blogPostID)
	return r.ExecutePaginatedQuery(query, pagination, &sorting)
}
