package mappers

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type BlogPostMapper struct{}

func NewBlogPostPersistenceMapper() *BlogPostMapper {
	return &BlogPostMapper{}
}

func (m *BlogPostMapper) MapToModel(entity entities.BlogPostEntity) models.BlogPostModel {
	return models.BlogPostModel{
		Id:               entity.ID.String(),
		Heading:          entity.Heading,
		PageTitle:        entity.PageTitle,
		Content:          entity.Content,
		ShortDescription: entity.ShortDescription,
		FeaturedImageUrl: entity.FeaturedImageUrl,
		UrlHandle:        entity.UrlHandle,
		AuthorName:       entity.Author,
		Visible:          entity.Visible,
		CreatedAt:        entity.CreatedAt.Format(time.RFC3339),
	}
}

func (m *BlogPostMapper) MapToEntity(model models.BlogPostModel) entities.BlogPostEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}
	var createdAt time.Time
	if model.CreatedAt == "" {
		createdAt = time.Now()
	} else {
		createdAt, _ = time.Parse(time.RFC3339, model.CreatedAt)
	}
	return entities.BlogPostEntity{
		ID:               id,
		Heading:          model.Heading,
		PageTitle:        model.PageTitle,
		Content:          model.Content,
		ShortDescription: model.ShortDescription,
		FeaturedImageUrl: model.FeaturedImageUrl,
		UrlHandle:        model.UrlHandle,
		Author:           model.AuthorName,
		Visible:          model.Visible,
		CreatedAt:        createdAt,
		UpdatedAt:        time.Now(),
	}
}
