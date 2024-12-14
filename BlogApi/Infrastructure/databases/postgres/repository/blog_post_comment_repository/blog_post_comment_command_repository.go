package repository

import (
	"context"
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"

	"gorm.io/gorm"
)

type BlogPostCommentCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel]
}

func NewBlogPostCommentCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
) repository_interfaces.BlogPostCommentCommandRepository {
	return &BlogPostCommentCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel](postgresDatabase.DB, mapper),
	}
}

func (r *BlogPostCommentCommandRepositoryImpl) AddComment(ctx context.Context, comment *models.BlogPostCommentModel) (string, error) {
	var id string
	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		createdComment, err := r.Add(ctx, *comment)
		if err != nil {
			return err
		}
		id = createdComment.Id
		return nil
	})
	return id, err
}

func (r *BlogPostCommentCommandRepositoryImpl) UpdateComment(ctx context.Context, comment *models.BlogPostCommentModel) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		_, err := r.Update(ctx, *comment)
		return err
	})
}

func (r *BlogPostCommentCommandRepositoryImpl) RemoveComment(ctx context.Context, commentID string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		return r.Delete(ctx, commentID)
	})
}
