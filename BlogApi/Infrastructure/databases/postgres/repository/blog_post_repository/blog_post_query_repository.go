package repository

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	property_selector "identity-api/Domain/property_selector"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	"identity-api/Domain/sorting"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
)

type BlogPostQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostEntity, string, models.BlogPostModel]
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostModel]
}

func NewBlogPostQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostModel],
) repository_interfaces.BlogPostQueryRepository {
	return &BlogPostQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostEntity, string, models.BlogPostModel](postgresDatabase.DB, mapper),
		mapper:          mapper,
	}
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPostByID(id string) (*models.BlogPostModel, error) {
	return r.GetById(id)
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPostByUrlHandle(urlHandle string) (*models.BlogPostModel, error) {
	urlHandleProp := property_selector.NewPropertySelector[entities.BlogPostEntity]("url_handle")
	return r.GetFirstByProp(urlHandleProp, urlHandle)
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPosts(
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
) (*pagination.PaginatedResult[models.BlogPostModel], error) {
	return r.GetAll(pagination, sorting)
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPostsByAuthorID(
	authorID string,
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
) (*pagination.PaginatedResult[models.BlogPostModel], error) {
	authorIDProp := property_selector.NewPropertySelector[entities.BlogPostEntity]("user_id")
	return r.GetAllByProp(authorIDProp, authorID, pagination, sorting)
}

func (r *BlogPostQueryRepositoryImpl) GetVisibleBlogPosts(
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
) (*pagination.PaginatedResult[models.BlogPostModel], error) {
	visibleProp := property_selector.NewPropertySelector[entities.BlogPostEntity]("visible")
	return r.GetAllByProp(visibleProp, true, pagination, sorting)
}
