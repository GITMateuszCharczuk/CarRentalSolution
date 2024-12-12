package mappers

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type BlogPostLikeMapper struct{}

func NewBlogPostLikePersistenceMapper() *BlogPostLikeMapper {
	return &BlogPostLikeMapper{}
}

func (m *BlogPostLikeMapper) MapToModel(entity entities.BlogPostLikeEntity) models.BlogPostLikeModel {
	return models.BlogPostLikeModel{
		Id:         entity.ID.String(),
		BlogPostId: entity.BlogPostID.String(),
		UserId:     entity.UserID.String(),
		CreatedAt:  entity.CreatedAt.Format(time.RFC3339),
	}
}

func (m *BlogPostLikeMapper) MapToEntity(model models.BlogPostLikeModel) entities.BlogPostLikeEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}

	blogPostID, _ := uuid.Parse(model.BlogPostId)
	userID, _ := uuid.Parse(model.UserId)

	return entities.BlogPostLikeEntity{
		ID:         id,
		BlogPostID: blogPostID,
		UserID:     userID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
