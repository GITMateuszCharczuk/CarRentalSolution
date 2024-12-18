package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_comment_repository"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

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
