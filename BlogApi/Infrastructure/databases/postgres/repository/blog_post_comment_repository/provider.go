package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

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
	ProvideBlogPostCommentQueryRepository,
	ProvideBlogPostCommentCommandRepository,
)
