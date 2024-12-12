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
	return r.GetById(id)
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagByName(name string) (*models.TagModel, error) {
	nameProp := property_selector.NewPropertySelector[entities.BlogPostTagEntity]("name")
	return r.GetFirstByProp(nameProp, name)
}

func (r *BlogPostTagQueryRepositoryImpl) GetAllTags(
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
) (*pagination.PaginatedResult[models.TagModel], error) {
	return r.GetAll(pagination, sorting)
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagsByBlogPostID(blogPostID string) ([]models.TagModel, error) {
	var entities []entities.BlogPostTagEntity
	if err := r.dbContext.Joins("JOIN blog_post_tags ON blog_post_tags.tag_id = blog_post_tag_entities.id").
		Where("blog_post_tags.blog_post_id = ?", blogPostID).
		Find(&entities).Error; err != nil {
		return nil, err
	}

	models := make([]models.TagModel, len(entities))
	for i, entity := range entities {
		models[i] = r.mapper.MapToModel(entity)
	}
	return models, nil
}

func (r *BlogPostTagQueryRepositoryImpl) GetBlogPostsByTag(
	tagName string,
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
) (*pagination.PaginatedResult[models.BlogPostModel], error) {
	var blogPostEntities []entities.BlogPostEntity
	var total int64

	query := r.dbContext.Model(&entities.BlogPostEntity{}).
		Joins("JOIN blog_post_tags ON blog_post_tags.blog_post_id = blog_post_entities.id").
		Joins("JOIN blog_post_tag_entities ON blog_post_tag_entities.id = blog_post_tags.tag_id").
		Where("blog_post_tag_entities.name = ?", tagName)

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := query.Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order(sorting.GetOrderBy()).
		Find(&blogPostEntities).Error; err != nil {
		return nil, err
	}

	blogPostModels := make([]models.BlogPostModel, len(blogPostEntities))
	for i, entity := range blogPostEntities {
		blogPostModels[i] = r.mapper.MapToModel(entity)
	}

	return pagination.CreatePaginatedResult(blogPostModels, total), nil
}
