package repository

import (
	"context"
	"errors"
	models "identity-api/Domain/models/domestic"
	tag_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_tag_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPostTagCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostTagEntity, string, models.TagModel]
}

func NewBlogPostTagCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.TagModel],
) tag_repository_interfaces.BlogPostTagCommandRepository {
	return &BlogPostTagCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.BlogPostTagEntity, string, models.TagModel](postgresDatabase.DB, mapper),
	}
}

func (r *BlogPostTagCommandRepositoryImpl) AddTagToBlogPost(ctx context.Context, tag *models.TagModel, blogPostEntity entities.BlogPostEntity) (*models.TagModel, error) {
	var resultTag *models.TagModel

	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		tagEntity := entities.BlogPostTagEntity{
			ID:        uuid.New(),
			Name:      tag.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := tx.Create(&tagEntity).Error; err != nil {
			if !errors.Is(err, gorm.ErrDuplicatedKey) {
				return err
			}
			if err := tx.Where("name = ?", tagEntity.Name).First(&tagEntity).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&blogPostEntity).Association("Tags").Append(&tagEntity); err != nil {
			return err
		}

		resultTag = &models.TagModel{
			Id:   tagEntity.ID.String(),
			Name: tagEntity.Name,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resultTag, nil
}

func (r *BlogPostTagCommandRepositoryImpl) AddTagsToBlogPost(ctx context.Context, blogPostEntity entities.BlogPostEntity, tagNames []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		for _, tagName := range tagNames {
			_, err := r.AddTagToBlogPost(ctx, &models.TagModel{Name: tagName}, blogPostEntity)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *BlogPostTagCommandRepositoryImpl) CleanupUnusedTags(ctx context.Context) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		return tx.Exec(`
            DELETE FROM blog_post_tag_entities 
            WHERE id IN (
                SELECT t.id 
                FROM blog_post_tag_entities t 
                LEFT JOIN blog_post_tags bt ON t.id = bt.tag_id 
                WHERE bt.blog_post_id IS NULL
            )
        `).Error
	})
}

func (r *BlogPostTagCommandRepositoryImpl) ModifyTagsForBlogPost(ctx context.Context, blogPostEntity entities.BlogPostEntity, newTagNames []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var existingTags []entities.BlogPostTagEntity
		if err := tx.Model(&blogPostEntity).Association("Tags").Find(&existingTags); err != nil {
			return err
		}

		existingTagNames := make(map[string]entities.BlogPostTagEntity)
		for _, tag := range existingTags {
			existingTagNames[tag.Name] = tag
		}

		newTagNamesMap := make(map[string]bool)
		for _, name := range newTagNames {
			newTagNamesMap[name] = true
		}

		for _, existingTag := range existingTags {
			if !newTagNamesMap[existingTag.Name] {
				if err := tx.Model(&blogPostEntity).Association("Tags").Delete(&existingTag); err != nil {
					return err
				}
			}
		}

		for _, newTagName := range newTagNames {
			if _, exists := existingTagNames[newTagName]; !exists {
				var tag entities.BlogPostTagEntity

				result := tx.Where("name = ?", newTagName).FirstOrCreate(&tag, entities.BlogPostTagEntity{
					ID:        uuid.New(),
					Name:      newTagName,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
				if result.Error != nil {
					return result.Error
				}

				if err := tx.Model(&blogPostEntity).Association("Tags").Append(&tag); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
