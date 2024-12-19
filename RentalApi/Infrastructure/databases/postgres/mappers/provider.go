package mappers

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Infrastructure/databases/postgres/entities"
	base "rental-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

func ProvideCarOfferPersistenceMapper() base.PersistenceMapper[entities.CarOfferEntity, models.CarOfferModel] {
	return NewCarOfferPersistenceMapper()
}

func ProvideCarOrderPersistenceMapper() base.PersistenceMapper[entities.CarOrderEntity, models.CarOrderModel] {
	return NewCarOrderPersistenceMapper()
}

func ProvideCarTagPersistenceMapper() base.PersistenceMapper[entities.CarTagEntity, models.CarOfferTagModel] {
	return NewCarTagPersistenceMapper()
}

func ProvideCarImagePersistenceMapper() base.PersistenceMapper[entities.CarImageEntity, models.CarOfferImageModel] {
	return NewCarImagePersistenceMapper()
}

var WireSet = wire.NewSet(
	ProvideCarOfferPersistenceMapper,
	ProvideCarOrderPersistenceMapper,
	ProvideCarTagPersistenceMapper,
	ProvideCarImagePersistenceMapper,
)
