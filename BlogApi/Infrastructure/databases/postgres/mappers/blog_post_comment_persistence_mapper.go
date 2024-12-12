package mappers

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type BlogPostCommentMapper struct{}

func NewBlogPostCommentPersistenceMapper() *BlogPostCommentMapper {
	return &BlogPostCommentMapper{}
}

func (m *BlogPostCommentMapper) MapToModel(entity entities.BlogPostCommentEntity) models.BlogPostCommentModel {
	return models.BlogPostCommentModel{
		Id:          entity.ID.String(),
		BlogPostId:  entity.BlogPostID.String(),
		UserId:      entity.UserID.String(),
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   entity.UpdatedAt.Format(time.RFC3339),
	}
}

func (m *BlogPostCommentMapper) MapToEntity(model models.BlogPostCommentModel) entities.BlogPostCommentEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}

	blogPostID, _ := uuid.Parse(model.BlogPostId)
	userID, _ := uuid.Parse(model.UserId)

	return entities.BlogPostCommentEntity{
		ID:          id,
		BlogPostID:  blogPostID,
		UserID:      userID,
		Description: model.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
