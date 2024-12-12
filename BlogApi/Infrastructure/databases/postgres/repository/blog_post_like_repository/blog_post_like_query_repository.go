package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
	"identity-api/Infrastructure/databases/postgres/repository/base/helpers"
)

type BlogPostLikeQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel]
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel]
}

func NewBlogPostLikeQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
) repository_interfaces.BlogPostLikeQueryRepository {
	return &BlogPostLikeQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel](postgresDatabase.DB, mapper),
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
	query = r.EnrichQuery(query, queryRecord, queryRecord2)
	return r.GetTotalCount(query)
}
