package repository

import (
	"context"
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_image_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"gorm.io/gorm"
)

type CarImageCommandRepositoryImpl struct {
	*base.CommandRepository[entities.CarImageEntity, string, models.CarOfferImageModel]
}

func NewCarImageCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarImageEntity, models.CarOfferImageModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarImageCommandRepository {
	return &CarImageCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.CarImageEntity, string, models.CarOfferImageModel](postgresDatabase.DB, mapper, uow),
	}
}

func (r *CarImageCommandRepositoryImpl) AddImageToCarOffer(ctx context.Context, carOfferId string, imageId string) (*string, error) {
	var result string
	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		image := models.CarOfferImageModel{
			CarOfferId: carOfferId,
			ImageId:    imageId,
		}
		createdImage, err := r.Add(ctx, image)
		if err != nil {
			return err
		}
		result = createdImage.Id
		return nil
	})
	return &result, err
}

func (r *CarImageCommandRepositoryImpl) AddImagesToCarOffer(ctx context.Context, carOfferId string, imageIds []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		for _, imageId := range imageIds {
			_, err := r.AddImageToCarOffer(ctx, carOfferId, imageId)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *CarImageCommandRepositoryImpl) ModifyImagesForCarOffer(ctx context.Context, carOfferId string, imageIds []string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var carOfferEntity entities.CarOfferEntity
		if err := tx.First(&carOfferEntity, "id = ?", carOfferId).Error; err != nil {
			return err
		}

		if err := tx.Where("car_offer_id = ?", carOfferId).Delete(&entities.CarImageEntity{}).Error; err != nil {
			return err
		}

		return r.AddImagesToCarOffer(ctx, carOfferId, imageIds)
	})
}

func (r *CarImageCommandRepositoryImpl) DeleteImageFromCarOffer(ctx context.Context, carOfferId string, imageId string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var entity entities.CarImageEntity
		if err := tx.Where("car_offer_id = ? AND image_id = ?", carOfferId, imageId).First(&entity).Error; err != nil {
			return err
		}
		return r.Delete(ctx, entity.ID.String())
	})
}
