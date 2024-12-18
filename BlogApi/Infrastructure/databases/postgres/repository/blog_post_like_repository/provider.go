package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_like_repository"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideBlogPostLikeQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostLikeQueryRepository {
	return NewBlogPostLikeQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideBlogPostLikeCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
	uow unit_of_work.UnitOfWork,
	blogPostLikeQueryRepository repository_interfaces.BlogPostLikeQueryRepository,
) repository_interfaces.BlogPostLikeCommandRepository {
	return NewBlogPostLikeCommandRepositoryImpl(postgresDatabase, mapper, blogPostLikeQueryRepository, uow)
}

var WireSet = wire.NewSet(
	ProvideBlogPostLikeQueryRepository,
	ProvideBlogPostLikeCommandRepository,
)
