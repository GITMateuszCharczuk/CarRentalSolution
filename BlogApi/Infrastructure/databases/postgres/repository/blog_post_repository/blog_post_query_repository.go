package repository

import (
	models "blog-api/Domain/models/domestic"
	"blog-api/Domain/pagination"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_repository"
	"blog-api/Domain/sorting"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	base "blog-api/Infrastructure/databases/postgres/repository/base"
	"blog-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
)

type BlogPostQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostEntity, string, models.BlogPostResponseModel]
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostResponseModel]
}

func NewBlogPostQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostResponseModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostQueryRepository {
	return &BlogPostQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostEntity, string, models.BlogPostResponseModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPostByID(id string) (*models.BlogPostResponseModel, error) {
	return r.GetById(id)
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPostAuthorId(id string) (*string, error) {
	var blogPost entities.BlogPostEntity
	if err := r.ConstructBaseQuery().Where("id = ?", id).First(&blogPost).Error; err != nil {
		return nil, err
	}
	res := blogPost.UserId.String()
	return &res, nil
}

func (r *BlogPostQueryRepositoryImpl) GetBlogPosts(
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
	ids []string,
	dateTimeFrom string,
	dateTimeTo string,
	authorIds []string,
	tagNames []string,
	visible string,
) (*pagination.PaginatedResult[models.BlogPostResponseModel], error) {
	query := r.ConstructBaseQuery()
	if len(tagNames) > 0 {
		query = query.Preload("Tags").
			Joins("JOIN blog_post_tags ON blog_post_tags.blog_post_id = blog_post_entities.id").
			Joins("JOIN blog_post_tag_entities ON blog_post_tag_entities.id = blog_post_tags.tag_id").
			Where("blog_post_tag_entities.name IN ?", tagNames)
	}
	queryRecords := []helpers.QueryRecord[entities.BlogPostEntity]{
		helpers.NewQueryRecord[entities.BlogPostEntity]("id", ids, "blog_post_entities"),
		helpers.NewQueryRecord[entities.BlogPostEntity]("user_id", authorIds, "blog_post_entities"),
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
