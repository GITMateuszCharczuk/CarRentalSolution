package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "identity-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideBlogPostCommentQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostCommentQueryRepository {
	return NewBlogPostCommentQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideBlogPostCommentCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostCommentCommandRepository {
	return NewBlogPostCommentCommandRepositoryImpl(postgresDatabase, mapper, uow)
}

var WireSet = wire.NewSet(
	ProvideBlogPostCommentQueryRepository,
	ProvideBlogPostCommentCommandRepository,
)
