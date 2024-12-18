package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_like_repository"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	base "blog-api/Infrastructure/databases/postgres/repository/base"
	"blog-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
)

type BlogPostLikeQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel]
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel]
}

func NewBlogPostLikeQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostLikeQueryRepository {
	return &BlogPostLikeQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *BlogPostLikeQueryRepositoryImpl) GetLikesCount(
	blogPostID string,
	userID string,
) (int64, error) {
	query := r.ConstructBaseQuery()
	queryRecord := helpers.NewQueryRecord[entities.BlogPostLikeEntity]("blog_post_id", blogPostID)
	queryRecord2 := helpers.NewQueryRecord[entities.BlogPostLikeEntity]("user_id", userID)
	query = r.ApplyWhereConditions(query, queryRecord, queryRecord2)
	return r.GetTotalCount(query)
}
