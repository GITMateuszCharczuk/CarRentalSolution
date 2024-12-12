package repository

import (
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"

	"github.com/google/uuid"
)

type BlogPostTagCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostTagEntity, string, models.TagModel]
}

func NewBlogPostTagCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.TagModel],
) repository_interfaces.BlogPostTagCommandRepository {
	return &BlogPostTagCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.BlogPostTagEntity, string, models.TagModel](postgresDatabase.DB, mapper),
	}
}

func (r *BlogPostTagCommandRepositoryImpl) CreateTag(tag *models.TagModel) (*models.TagModel, error) {
	result, err := r.Add(*tag)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BlogPostTagCommandRepositoryImpl) UpdateTag(tag *models.TagModel) (*models.TagModel, error) {
	result, err := r.Update(*tag)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BlogPostTagCommandRepositoryImpl) DeleteTag(id string) error {
	return r.Delete(id)
}

func (r *BlogPostTagCommandRepositoryImpl) AddTagToBlogPost(blogPostID string, tagID string) error {
	blogPostUUID, err := uuid.Parse(blogPostID)
	if err != nil {
		return err
	}

	tagUUID, err := uuid.Parse(tagID)
	if err != nil {
		return err
	}

	return r.dbContext.Exec(
		"INSERT INTO blog_post_tags (blog_post_id, tag_id) VALUES (?, ?)",
		blogPostUUID,
		tagUUID,
	).Error
}

func (r *BlogPostTagCommandRepositoryImpl) RemoveTagFromBlogPost(blogPostID string, tagID string) error {
	blogPostUUID, err := uuid.Parse(blogPostID)
	if err != nil {
		return err
	}

	tagUUID, err := uuid.Parse(tagID)
	if err != nil {
		return err
	}

	return r.dbContext.Exec(
		"DELETE FROM blog_post_tags WHERE blog_post_id = ? AND tag_id = ?",
		blogPostUUID,
		tagUUID,
	).Error
}
