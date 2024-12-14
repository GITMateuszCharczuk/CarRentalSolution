package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

func ProvideBlogPostLikeQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
) repository_interfaces.BlogPostLikeQueryRepository {
	return NewBlogPostLikeQueryRepositoryImpl(postgresDatabase, mapper)
}

func ProvideBlogPostLikeCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
) repository_interfaces.BlogPostLikeCommandRepository {
	return NewBlogPostLikeCommandRepositoryImpl(postgresDatabase, mapper)
}

var WireSet = wire.NewSet(
	ProvideBlogPostLikeQueryRepository,
	ProvideBlogPostLikeCommandRepository,
)
