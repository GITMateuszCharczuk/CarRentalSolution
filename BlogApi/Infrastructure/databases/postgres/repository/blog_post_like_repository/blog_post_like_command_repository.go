package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
)

type BlogPostLikeCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel]
}

func NewBlogPostLikeCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
) repository_interfaces.BlogPostLikeCommandRepository {
	return &BlogPostLikeCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel](postgresDatabase.DB, mapper),
	}
}

func (r *BlogPostLikeCommandRepositoryImpl) AddLike(blogPostID string, userID string) error {
	like := models.BlogPostLikeModel{
		BlogPostId: blogPostID,
		UserId:     userID,
	}

	_, err := r.Add(like)
	return err
}

func (r *BlogPostLikeCommandRepositoryImpl) RemoveLike(blogPostID string, userID string) error {
	var entity entities.BlogPostLikeEntity
	if err := r.DbContext.Where("blog_post_id = ? AND user_id = ?", blogPostID, userID).First(&entity).Error; err != nil {
		return err
	}
	return r.Delete(entity.ID.String())
}
