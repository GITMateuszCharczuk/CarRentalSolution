package repository

import (
	models "identity-api/Domain/models/domestic"
	blog_post_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_tag_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "identity-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideBlogPostTagQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.BlogPostTagModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostTagQueryRepository {
	return NewBlogPostTagQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideBlogPostTagCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.BlogPostTagModel],
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostTagCommandRepository {
	return NewBlogPostTagCommandRepositoryImpl(postgresDatabase, blogPostQueryRepository, mapper, uow)
}

var WireSet = wire.NewSet(
	ProvideBlogPostTagQueryRepository,
	ProvideBlogPostTagCommandRepository,
)
