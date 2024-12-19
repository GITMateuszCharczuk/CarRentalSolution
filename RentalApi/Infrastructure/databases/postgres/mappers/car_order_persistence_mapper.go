package mappers

import (
	"rental-api/Domain/constants"
	models "rental-api/Domain/models/domestic"
	"rental-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type CarOrderMapper struct{}

func NewCarOrderPersistenceMapper() *CarOrderMapper {
	return &CarOrderMapper{}
}

func (m *CarOrderMapper) MapToModel(entity entities.CarOrderEntity) models.CarOrderModel {
	var carOfferId string
	if entity.CarOfferId != nil {
		carOfferId = entity.CarOfferId.String()
	} else {
		carOfferId = ""
	}
	return models.CarOrderModel{
		Id:           entity.ID.String(),
		UserId:       entity.UserId.String(),
		CarOfferId:   carOfferId,
		StartDate:    entity.StartDate.Format(time.RFC3339),
		EndDate:      entity.EndDate.Format(time.RFC3339),
		NumOfDrivers: entity.NumOfDrivers,
		TotalCost:    entity.TotalCost,
		Status:       string(entity.Status),
	}
}

func (m *CarOrderMapper) MapToEntity(model models.CarOrderModel) entities.CarOrderEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}

	userId, _ := uuid.Parse(model.UserId)
	carOfferId, _ := uuid.Parse(model.CarOfferId)
	startDate, _ := time.Parse(time.RFC3339, model.StartDate)
	endDate, _ := time.Parse(time.RFC3339, model.EndDate)

	return entities.CarOrderEntity{
		ID:           id,
		UserId:       userId,
		CarOfferId:   &carOfferId,
		StartDate:    startDate,
		EndDate:      endDate,
		NumOfDrivers: model.NumOfDrivers,
		TotalCost:    model.TotalCost,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Status:       constants.CarOrderStatus(model.Status),
	}
}
