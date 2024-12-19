package mappers

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Infrastructure/databases/postgres/entities"
	"time"

	"github.com/google/uuid"
)

type CarOfferMapper struct{}

func NewCarOfferPersistenceMapper() *CarOfferMapper {
	return &CarOfferMapper{}
}

func (m *CarOfferMapper) MapToModel(entity entities.CarOfferEntity) models.CarOfferModel {
	return models.CarOfferModel{
		Id:                 entity.ID.String(),
		Heading:            entity.Heading,
		ShortDescription:   entity.ShortDescription,
		UrlHandle:          entity.UrlHandle,
		Horsepower:         entity.Horsepower,
		YearOfProduction:   entity.YearOfProduction,
		EngineDetails:      entity.EngineDetails,
		DriveDetails:       entity.DriveDetails,
		GearboxDetails:     entity.GearboxDetails,
		PublishedDate:      entity.PublishedDate.Format(time.RFC3339),
		Visible:            entity.Visible,
		OneNormalDayPrice:  entity.OneNormalDayPrice,
		OneWeekendDayPrice: entity.OneWeekendDayPrice,
		OneWeekPrice:       entity.OneWeekPrice,
		OneMonthPrice:      entity.OneMonthPrice,
		CreatedAt:          entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          entity.UpdatedAt.Format(time.RFC3339),
		CustodianEmail:     entity.CustodianEmail,
	}
}

func (m *CarOfferMapper) MapToEntity(model models.CarOfferModel) entities.CarOfferEntity {
	var id uuid.UUID
	if model.Id == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.Id)
	}

	publishedDate, _ := time.Parse(time.RFC3339, model.PublishedDate)

	var createdAt time.Time
	if model.CreatedAt == "" {
		createdAt = time.Now()
	} else {
		createdAt, _ = time.Parse(time.RFC3339, model.CreatedAt)
	}

	return entities.CarOfferEntity{
		ID:                 id,
		Heading:            model.Heading,
		ShortDescription:   model.ShortDescription,
		UrlHandle:          model.UrlHandle,
		Horsepower:         model.Horsepower,
		YearOfProduction:   model.YearOfProduction,
		EngineDetails:      model.EngineDetails,
		DriveDetails:       model.DriveDetails,
		GearboxDetails:     model.GearboxDetails,
		PublishedDate:      publishedDate,
		Visible:            model.Visible,
		OneNormalDayPrice:  model.OneNormalDayPrice,
		OneWeekendDayPrice: model.OneWeekendDayPrice,
		OneWeekPrice:       model.OneWeekPrice,
		OneMonthPrice:      model.OneMonthPrice,
		CustodianId:        uuid.MustParse(model.CustodianId),
		CustodianEmail:     model.CustodianEmail,
		CreatedAt:          createdAt,
		UpdatedAt:          time.Now(),
	}
}
