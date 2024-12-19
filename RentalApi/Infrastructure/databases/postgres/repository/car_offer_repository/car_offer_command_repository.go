package repository

import (
	"context"
	"log"
	models "rental-api/Domain/models/domestic"
	car_image_repository "rental-api/Domain/repository_interfaces/car_image_repository"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	car_tag_repository "rental-api/Domain/repository_interfaces/car_tag_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"gorm.io/gorm"
)

type CarOfferCommandRepositoryImpl struct {
	*base.CommandRepository[entities.CarOfferEntity, string, models.CarOfferModel]
	mapper                    mappers.PersistenceMapper[entities.CarOfferEntity, models.CarOfferModel]
	carTagCommandRepository   car_tag_repository.CarTagCommandRepository
	carImageCommandRepository car_image_repository.CarImageCommandRepository
}

func NewCarOfferCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOfferEntity, models.CarOfferModel],
	uow unit_of_work.UnitOfWork,
	carTagCommandRepository car_tag_repository.CarTagCommandRepository,
	carImageCommandRepository car_image_repository.CarImageCommandRepository,
) repository_interfaces.CarOfferCommandRepository {
	return &CarOfferCommandRepositoryImpl{
		CommandRepository:         base.NewCommandRepository[entities.CarOfferEntity, string, models.CarOfferModel](postgresDatabase.DB, mapper, uow),
		carTagCommandRepository:   carTagCommandRepository,
		carImageCommandRepository: carImageCommandRepository,
		mapper:                    mapper,
	}
}

func (r *CarOfferCommandRepositoryImpl) CreateCarOffer(ctx context.Context, carOffer *models.CarOfferModel, tags []string, imageIds []string) (*string, error) {
	var result string
	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		createdOffer, err := r.Add(ctx, *carOffer)
		if err != nil {
			return err
		}

		if len(tags) > 0 {
			err := r.carTagCommandRepository.AddTagsToCarOffer(ctx, createdOffer.Id, tags)
			if err != nil {
				return err
			}
		}

		if len(imageIds) > 0 {
			err := r.carImageCommandRepository.AddImagesToCarOffer(ctx, createdOffer.Id, imageIds)
			if err != nil {
				return err
			}
		}

		result = createdOffer.Id
		return nil
	})
	return &result, err
}

func (r *CarOfferCommandRepositoryImpl) UpdateCarOffer(ctx context.Context, carOffer *models.CarOfferModel, tags []string, imageIds []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		entity := r.mapper.MapToEntity(*carOffer)

		if err := tx.Omit("Tags").Save(&entity).Error; err != nil {
			log.Println(err)
			return err
		}

		if len(tags) > 0 {
			if err := r.carTagCommandRepository.ModifyTagsForCarOffer(ctx, entity.ID.String(), tags); err != nil {
				log.Println(err)
				return err
			}
		}

		if len(imageIds) > 0 {
			if err := r.carImageCommandRepository.ModifyImagesForCarOffer(ctx, entity.ID.String(), imageIds); err != nil {
				log.Println(err)
				return err
			}
		}

		return nil
	})
}

func (r *CarOfferCommandRepositoryImpl) DeleteCarOffer(ctx context.Context, id string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var carOffer entities.CarOfferEntity
		if err := tx.First(&carOffer, "id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Model(&carOffer).Association("Tags").Clear(); err != nil {
			return err
		}

		if err := r.Delete(ctx, id); err != nil {
			return err
		}
		return r.carTagCommandRepository.CleanupUnusedTags(ctx)
	})
}
