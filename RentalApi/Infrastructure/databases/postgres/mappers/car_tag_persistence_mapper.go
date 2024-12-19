package mappers

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type CarTagMapper struct{}

func NewCarTagPersistenceMapper() *CarTagMapper {
	return &CarTagMapper{}
}

func (m *CarTagMapper) MapToModel(entity entities.CarTagEntity) models.CarOfferTagModel {
	return models.CarOfferTagModel{
		Id:   entity.ID.String(),
		Name: entity.Name,
	}
}

func (m *CarTagMapper) MapToEntity(model models.CarOfferTagModel) entities.CarTagEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}

	return entities.CarTagEntity{
		ID:        id,
		Name:      model.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
