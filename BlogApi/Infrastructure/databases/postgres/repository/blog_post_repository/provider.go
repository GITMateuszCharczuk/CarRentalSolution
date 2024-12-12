package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

func ProvideBlogPostQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostModel],
) repository_interfaces.BlogPostQueryRepository {
	return NewBlogPostQueryRepositoryImpl(postgresDatabase, mapper)
}

func ProvideBlogPostCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostModel],
) repository_interfaces.BlogPostCommandRepository {
	return NewBlogPostCommandRepositoryImpl(postgresDatabase, mapper)
}

func ProvideBlogPostTagQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.TagModel],
) repository_interfaces.BlogPostTagQueryRepository {
	return NewBlogPostTagQueryRepositoryImpl(postgresDatabase, mapper)
}

func ProvideBlogPostTagCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.TagModel],
) repository_interfaces.BlogPostTagCommandRepository {
	return NewBlogPostTagCommandRepositoryImpl(postgresDatabase, mapper)
}

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

func ProvideBlogPostCommentQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
) repository_interfaces.BlogPostCommentQueryRepository {
	return NewBlogPostCommentQueryRepositoryImpl(postgresDatabase, mapper)
}

func ProvideBlogPostCommentCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
) repository_interfaces.BlogPostCommentCommandRepository {
	return NewBlogPostCommentCommandRepositoryImpl(postgresDatabase, mapper)
}

var WireSet = wire.NewSet(
	ProvideBlogPostQueryRepository,
	ProvideBlogPostCommandRepository,
	ProvideBlogPostTagQueryRepository,
	ProvideBlogPostTagCommandRepository,
	ProvideBlogPostLikeQueryRepository,
	ProvideBlogPostLikeCommandRepository,
	ProvideBlogPostCommentQueryRepository,
	ProvideBlogPostCommentCommandRepository,
)
