package mappers

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type BlogPostTagMapper struct{}

func NewBlogPostTagPersistenceMapper() *BlogPostTagMapper {
	return &BlogPostTagMapper{}
}

func (m *BlogPostTagMapper) MapToModel(entity entities.BlogPostTagEntity) models.BlogPostTagModel {
	return models.BlogPostTagModel{
		Id:   entity.ID.String(),
		Name: entity.Name,
	}
}

func (m *BlogPostTagMapper) MapToEntity(model models.BlogPostTagModel) entities.BlogPostTagEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}

	return entities.BlogPostTagEntity{
		ID:        id,
		Name:      model.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
