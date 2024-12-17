package repository

import (
	"context"
	models "identity-api/Domain/models/domestic"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	tag_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_tag_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "identity-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	"log"

	"gorm.io/gorm"
)

type BlogPostCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostEntity, string, models.BlogPostRequestModel]
	mapper                       mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostRequestModel]
	blogPostTagCommandRepository tag_repository_interfaces.BlogPostTagCommandRepository
}

func NewBlogPostCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostEntity, models.BlogPostRequestModel],
	blogPostTagCommandRepository tag_repository_interfaces.BlogPostTagCommandRepository,
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostCommandRepository {
	return &BlogPostCommandRepositoryImpl{
		CommandRepository:            base.NewCommandRepository[entities.BlogPostEntity, string, models.BlogPostRequestModel](postgresDatabase.DB, mapper, uow),
		mapper:                       mapper,
		blogPostTagCommandRepository: blogPostTagCommandRepository,
	}
}

func (r *BlogPostCommandRepositoryImpl) CreateBlogPost(ctx context.Context, blogPost *models.BlogPostRequestModel) (*string, error) {
	var result string

	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		entity := r.mapper.MapToEntity(*blogPost)

		if err := tx.Create(&entity).Error; err != nil {
			return err
		}
		if len(blogPost.Tags) > 0 {
			err := r.blogPostTagCommandRepository.AddTagsToBlogPost(ctx, entity.ID.String(), blogPost.Tags)
			if err != nil {
				log.Println(err)
				return err
			}
		}

		result = entity.ID.String()
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *BlogPostCommandRepositoryImpl) UpdateBlogPost(ctx context.Context, blogPost *models.BlogPostRequestModel) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		entity := r.mapper.MapToEntity(*blogPost)

		if err := tx.Omit("Tags").Save(&entity).Error; err != nil {
			log.Println(err)
			return err
		}

		if len(blogPost.Tags) > 0 {
			if err := r.blogPostTagCommandRepository.ModifyTagsForBlogPost(ctx, entity.ID.String(), blogPost.Tags); err != nil {
				log.Println(err)
				return err
			}
		}

		return nil
	})
}

func (r *BlogPostCommandRepositoryImpl) DeleteBlogPost(ctx context.Context, id string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var blogPost entities.BlogPostEntity
		if err := tx.First(&blogPost, "id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Model(&blogPost).Association("Tags").Clear(); err != nil {
			return err
		}

		if err := r.Delete(ctx, id); err != nil {
			return err
		}
		return r.blogPostTagCommandRepository.CleanupUnusedTags(ctx)
	})
}
