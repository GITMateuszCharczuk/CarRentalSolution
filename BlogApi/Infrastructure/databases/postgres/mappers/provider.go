package mappers

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Infrastructure/databases/postgres/entities"
	base "identity-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

func ProvideBlogPostPersistenceMapper() base.PersistenceMapper[entities.BlogPostEntity, models.BlogPostModel] {
	return NewBlogPostPersistenceMapper()
}

func ProvideBlogPostTagPersistenceMapper() base.PersistenceMapper[entities.BlogPostTagEntity, models.TagModel] {
	return NewBlogPostTagPersistenceMapper()
}

func ProvideBlogPostLikePersistenceMapper() base.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel] {
	return NewBlogPostLikePersistenceMapper()
}

func ProvideBlogPostCommentPersistenceMapper() base.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel] {
	return NewBlogPostCommentPersistenceMapper()
}

var WireSet = wire.NewSet(
	ProvideBlogPostPersistenceMapper,
	ProvideBlogPostTagPersistenceMapper,
	ProvideBlogPostLikePersistenceMapper,
	ProvideBlogPostCommentPersistenceMapper,
)
