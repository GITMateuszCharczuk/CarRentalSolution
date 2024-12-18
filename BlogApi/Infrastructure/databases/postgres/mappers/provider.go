package mappers

import (
	models "blog-api/Domain/models/domestic"
	"blog-api/Infrastructure/databases/postgres/entities"
	base "blog-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

func ProvideBlogPostRequestPersistenceMapper() base.PersistenceMapper[entities.BlogPostEntity, models.BlogPostRequestModel] {
	return NewBlogPostRequestPersistenceMapper()
}

func ProvideBlogPostResponsePersistenceMapper() base.PersistenceMapper[entities.BlogPostEntity, models.BlogPostResponseModel] {
	return NewBlogPostResponsePersistenceMapper()
}

func ProvideBlogPostTagPersistenceMapper() base.PersistenceMapper[entities.BlogPostTagEntity, models.BlogPostTagModel] {
	return NewBlogPostTagPersistenceMapper()
}

func ProvideBlogPostLikePersistenceMapper() base.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel] {
	return NewBlogPostLikePersistenceMapper()
}

func ProvideBlogPostCommentPersistenceMapper() base.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel] {
	return NewBlogPostCommentPersistenceMapper()
}

var WireSet = wire.NewSet(
	ProvideBlogPostRequestPersistenceMapper,
	ProvideBlogPostResponsePersistenceMapper,
	ProvideBlogPostTagPersistenceMapper,
	ProvideBlogPostLikePersistenceMapper,
	ProvideBlogPostCommentPersistenceMapper,
)
