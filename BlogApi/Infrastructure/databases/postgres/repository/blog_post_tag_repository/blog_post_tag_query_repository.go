package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_tag_repository"
	"blog-api/Domain/sorting"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	base "blog-api/Infrastructure/databases/postgres/repository/base"
	"blog-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
)

type BlogPostTagQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostTagEntity, string, models.BlogPostTagModel]
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.BlogPostTagModel]
}

func NewBlogPostTagQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.BlogPostTagModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostTagQueryRepository {
	return &BlogPostTagQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostTagEntity, string, models.BlogPostTagModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagByID(id string) (*models.BlogPostTagModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.BlogPostTagEntity]("id", id)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagByName(name string) (*models.BlogPostTagModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.BlogPostTagEntity]("name", name)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *BlogPostTagQueryRepositoryImpl) GetTagsByBlogPostID(
	blogPostID string,
	sorting sorting.Sortable,
) (*[]models.BlogPostTagModel, error) {
	db := r.GetUnitOfWork().GetTransaction().Model(&entities.BlogPostTagEntity{})
	if blogPostID != "" {
		db = db.Joins("JOIN blog_post_tags ON blog_post_tags.tag_id = blog_post_tag_entities.id").
			Where("blog_post_tags.blog_post_id = ?", blogPostID)
	}
	return r.ExecuteSortedQuery(db, &sorting)
}
