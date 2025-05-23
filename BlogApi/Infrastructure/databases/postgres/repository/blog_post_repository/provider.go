package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_repository"
	tag_repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_tag_repository"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideBlogPostQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostResponseModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostQueryRepository {
	return NewBlogPostQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideBlogPostCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostRequestModel],
	blogPostTagCommandRepository tag_repository_interfaces.BlogPostTagCommandRepository,
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostCommandRepository {
	return NewBlogPostCommandRepositoryImpl(postgresDatabase, mapper, blogPostTagCommandRepository, uow)
}

var WireSet = wire.NewSet(
	ProvideBlogPostQueryRepository,
	ProvideBlogPostCommandRepository,
)
