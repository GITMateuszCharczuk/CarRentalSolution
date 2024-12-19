package repository

import (
	"context"
	models "rental-api/Domain/models/domestic"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_tag_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CarTagCommandRepositoryImpl struct {
	*base.CommandRepository[entities.CarTagEntity, string, models.CarOfferTagModel]
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository
}

func NewCarTagCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarTagEntity, models.CarOfferTagModel],
	uow unit_of_work.UnitOfWork,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
) repository_interfaces.CarTagCommandRepository {
	return &CarTagCommandRepositoryImpl{
		CommandRepository:       base.NewCommandRepository[entities.CarTagEntity, string, models.CarOfferTagModel](postgresDatabase.DB, mapper, uow),
		carOfferQueryRepository: carOfferQueryRepository,
	}
}

func (r *CarTagCommandRepositoryImpl) AddTagToCarOffer(ctx context.Context, tag *models.CarOfferTagModel, carOfferEntity entities.CarOfferEntity) (*models.CarOfferTagModel, error) {
	var resultTag *models.CarOfferTagModel

	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var tagEntity entities.CarTagEntity

		condition := entities.CarTagEntity{Name: tag.Name}
		defaults := entities.CarTagEntity{
			ID:   uuid.New(),
			Name: tag.Name,
		}

		if err := tx.Where(condition).Attrs(defaults).FirstOrCreate(&tagEntity).Error; err != nil {
			return err
		}

		if err := tx.Model(&carOfferEntity).Association("Tags").Append(&tagEntity); err != nil {
			return err
		}

		resultTag = &models.CarOfferTagModel{
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

func (r *CarTagCommandRepositoryImpl) AddTagsToCarOffer(ctx context.Context, carOfferId string, tagNames []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var carOfferEntity entities.CarOfferEntity
		if err := tx.First(&carOfferEntity, "id = ?", carOfferId).Error; err != nil {
			return err
		}

		for _, tagName := range tagNames {
			_, err := r.AddTagToCarOffer(ctx, &models.CarOfferTagModel{Name: tagName}, carOfferEntity)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *CarTagCommandRepositoryImpl) ModifyTagsForCarOffer(ctx context.Context, carOfferId string, newTagNames []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var carOfferEntity entities.CarOfferEntity
		if err := tx.First(&carOfferEntity, "id = ?", carOfferId).Error; err != nil {
			return err
		}

		if err := tx.Model(&carOfferEntity).Association("Tags").Clear(); err != nil {
			return err
		}

		return r.AddTagsToCarOffer(ctx, carOfferId, newTagNames)
	})
}

func (r *CarTagCommandRepositoryImpl) CleanupUnusedTags(ctx context.Context) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var unusedTags []entities.CarTagEntity

		if err := tx.Model(&entities.CarTagEntity{}).
			Joins("LEFT JOIN car_offer_tags ON car_offer_tags.tag_id = car_tag_entities.id").
			Where("car_offer_tags.car_offer_id IS NULL").
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
