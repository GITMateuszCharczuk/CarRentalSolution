package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_tag_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

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

var WireSet = wire.NewSet(
	ProvideBlogPostTagQueryRepository,
	ProvideBlogPostTagCommandRepository,
)
