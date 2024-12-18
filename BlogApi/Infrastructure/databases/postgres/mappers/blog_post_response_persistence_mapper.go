package mappers

import (
	models "blog-api/Domain/models/domestic"
	"blog-api/Infrastructure/databases/postgres/entities"
	"time"
)

type BlogPostResponseMapper struct{}

func NewBlogPostResponsePersistenceMapper() *BlogPostResponseMapper {
	return &BlogPostResponseMapper{}
}

func (m *BlogPostResponseMapper) MapToModel(entity entities.BlogPostEntity) models.BlogPostResponseModel {
	return models.BlogPostResponseModel{
		Id:               entity.ID.String(),
		Heading:          entity.Heading,
		PageTitle:        entity.PageTitle,
		Content:          entity.Content,
		ShortDescription: entity.ShortDescription,
		FeaturedImageUrl: entity.FeaturedImageUrl,
		UrlHandle:        entity.UrlHandle,
		AuthorName:       entity.Author,
		PublishedDate:    entity.PublishedDate.Format(time.RFC3339),
		CreatedAt:        entity.CreatedAt.Format(time.RFC3339),
	}
}

func (m *BlogPostResponseMapper) MapToEntity(model models.BlogPostResponseModel) entities.BlogPostEntity {
	panic("Not intended to be used")
}
