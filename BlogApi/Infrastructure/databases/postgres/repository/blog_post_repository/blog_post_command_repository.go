package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
)

type BlogPostCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostEntity, string, models.BlogPostModel]
}

func NewBlogPostCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostModel],
) repository_interfaces.BlogPostCommandRepository {
	return &BlogPostCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.BlogPostEntity, string, models.BlogPostModel](postgresDatabase.DB, mapper),
	}
}

func (r *BlogPostCommandRepositoryImpl) CreateBlogPost(blogPost *models.BlogPostModel) (*models.BlogPostModel, error) {
	result, err := r.Add(*blogPost)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BlogPostCommandRepositoryImpl) UpdateBlogPost(blogPost *models.BlogPostModel) (*models.BlogPostModel, error) {
	result, err := r.Update(*blogPost)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BlogPostCommandRepositoryImpl) DeleteBlogPost(id string) error {
	return r.Delete(id)
}
