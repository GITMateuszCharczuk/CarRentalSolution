package mappers

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type CarImageMapper struct{}

func NewCarImagePersistenceMapper() *CarImageMapper {
	return &CarImageMapper{}
}

func (m *CarImageMapper) MapToModel(entity entities.CarImageEntity) models.CarOfferImageModel {
	return models.CarOfferImageModel{
		Id:         entity.ID.String(),
		ImageId:    entity.ImageId,
		CarOfferId: entity.CarOfferId.String(),
		CreatedAt:  entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  entity.UpdatedAt.Format(time.RFC3339),
	}
}

func (m *CarImageMapper) MapToEntity(model models.CarOfferImageModel) entities.CarImageEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}

	carOfferId, _ := uuid.Parse(model.CarOfferId)

	return entities.CarImageEntity{
		ID:         id,
		ImageId:    model.ImageId,
		CarOfferId: carOfferId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
