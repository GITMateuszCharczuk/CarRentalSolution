package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_like_repository"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	base "blog-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	"context"

	"gorm.io/gorm"
)

type BlogPostLikeCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel]
	QueryRepository repository_interfaces.BlogPostLikeQueryRepository
}

func NewBlogPostLikeCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostLikeEntity, models.BlogPostLikeModel],
	blogPostLikeQueryRepository repository_interfaces.BlogPostLikeQueryRepository,
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostLikeCommandRepository {
	return &BlogPostLikeCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.BlogPostLikeEntity, string, models.BlogPostLikeModel](postgresDatabase.DB, mapper, uow),
		QueryRepository:   blogPostLikeQueryRepository,
	}
}

func (r *BlogPostLikeCommandRepositoryImpl) AddLike(ctx context.Context, blogPostID string, userID string) (*string, error) {
	var result string
	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		count, err := r.QueryRepository.GetLikesCount(blogPostID, userID)
		if count > 0 {
			result = ""
			return nil
		}
		like := models.BlogPostLikeModel{
			BlogPostId: blogPostID,
			UserId:     userID,
		}
		like, err = r.Add(ctx, like)
		if err != nil {
			return err
		}
		result = like.Id
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
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
