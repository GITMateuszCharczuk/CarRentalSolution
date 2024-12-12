package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
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

func (r *BlogPostCommentCommandRepositoryImpl) AddComment(comment *models.BlogPostCommentModel) (string, error) {
	createdComment, err := r.Add(*comment)
	return createdComment.Id, err
}

func (r *BlogPostCommentCommandRepositoryImpl) UpdateComment(comment *models.BlogPostCommentModel) error {
	_, err := r.Update(*comment)
	return err
}

func (r *BlogPostCommentCommandRepositoryImpl) RemoveComment(commentID string) error {
	return r.Delete(commentID)
}
