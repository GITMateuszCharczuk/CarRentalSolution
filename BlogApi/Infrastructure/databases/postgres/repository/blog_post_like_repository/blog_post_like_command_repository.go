package repository

import (
	"context"
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"

	"gorm.io/gorm"
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

func (r *BlogPostLikeCommandRepositoryImpl) AddLike(ctx context.Context, blogPostID string, userID string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		like := models.BlogPostLikeModel{
			BlogPostId: blogPostID,
			UserId:     userID,
		}
		_, err := r.Add(ctx, like)
		return err
	})
}

func (r *BlogPostLikeCommandRepositoryImpl) RemoveLike(ctx context.Context, blogPostID string, userID string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var entity entities.BlogPostLikeEntity
		if err := tx.Where("blog_post_id = ? AND user_id = ?", blogPostID, userID).First(&entity).Error; err != nil {
			return err
		}
		return r.Delete(ctx, entity.ID.String())
	})
}
