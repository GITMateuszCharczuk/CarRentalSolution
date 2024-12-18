package repository

import (
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_repository"
	tag_repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_tag_repository"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	base "blog-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPostTagCommandRepositoryImpl struct {
	*base.CommandRepository[entities.BlogPostTagEntity, string, models.BlogPostTagModel]
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository
}

func NewBlogPostTagCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository,
	mapper mappers.PersistenceMapper[entities.BlogPostTagEntity, models.BlogPostTagModel],
	uow unit_of_work.UnitOfWork,
) tag_repository_interfaces.BlogPostTagCommandRepository {
	return &BlogPostTagCommandRepositoryImpl{
		CommandRepository:       base.NewCommandRepository[entities.BlogPostTagEntity, string, models.BlogPostTagModel](postgresDatabase.DB, mapper, uow),
		blogPostQueryRepository: blogPostQueryRepository,
	}
}

func (r *BlogPostTagCommandRepositoryImpl) AddTagToBlogPost(ctx context.Context, tag *models.BlogPostTagModel, blogPostEntity entities.BlogPostEntity) (*models.BlogPostTagModel, error) {
	var resultTag *models.BlogPostTagModel

	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var tagEntity entities.BlogPostTagEntity

		condition := entities.BlogPostTagEntity{Name: tag.Name}

		defaults := entities.BlogPostTagEntity{
			ID:   uuid.New(),
			Name: tag.Name,
		}

		if err := tx.Where(condition).Attrs(defaults).FirstOrCreate(&tagEntity).Error; err != nil {
			return err
		}

		if err := tx.Model(&blogPostEntity).Association("Tags").Append(&tagEntity); err != nil {
			return err
		}

		resultTag = &models.BlogPostTagModel{
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

func (r *BlogPostTagCommandRepositoryImpl) AddTagsToBlogPost(ctx context.Context, blogPostId string, tagNames []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var blogPostEntity entities.BlogPostEntity
		if err := tx.First(&blogPostEntity, "id = ?", blogPostId).Error; err != nil {
			return err
		}
		for _, tagName := range tagNames {
			_, err := r.AddTagToBlogPost(ctx, &models.BlogPostTagModel{Name: tagName}, blogPostEntity)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *BlogPostTagCommandRepositoryImpl) CleanupUnusedTags(ctx context.Context) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var unusedTags []entities.BlogPostTagEntity

		if err := tx.Model(&entities.BlogPostTagEntity{}).
			Joins("LEFT JOIN blog_post_tags ON blog_post_tags.tag_id = blog_post_tag_entities.id").
			Where("blog_post_tags.blog_post_id IS NULL").
			Find(&unusedTags).Error; err != nil {
			return err
		}

		if len(unusedTags) > 0 {
			if err := tx.Delete(&unusedTags).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *BlogPostTagCommandRepositoryImpl) ModifyTagsForBlogPost(ctx context.Context, blogPostId string, newTagNames []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var blogPostEntity entities.BlogPostEntity
		if err := tx.First(&blogPostEntity, "id = ?", blogPostId).Error; err != nil {
			return err
		}

		if err := tx.Model(&blogPostEntity).Association("Tags").Clear(); err != nil {
			return err
		}

		return r.AddTagsToBlogPost(ctx, blogPostId, newTagNames)
	})
}
