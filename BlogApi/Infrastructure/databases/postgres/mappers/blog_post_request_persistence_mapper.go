package mappers

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type BlogPostRequestMapper struct{}

func NewBlogPostRequestPersistenceMapper() *BlogPostRequestMapper {
	return &BlogPostRequestMapper{}
}

func (m *BlogPostRequestMapper) MapToModel(entity entities.BlogPostEntity) models.BlogPostRequestModel {
	panic("Not intended to be used")
}

func (m *BlogPostRequestMapper) MapToEntity(model models.BlogPostRequestModel) entities.BlogPostEntity {
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
