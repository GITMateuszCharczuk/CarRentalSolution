package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_comment_repository"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	base "blog-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	"context"

	"gorm.io/gorm"
)

type BlogPostCommentCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel]
}

func NewBlogPostCommentCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostCommentCommandRepository {
	return &BlogPostCommentCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel](postgresDatabase.DB, mapper, uow),
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
