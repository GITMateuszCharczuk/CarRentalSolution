package mappers

import (
	models "identity-api/Domain/models/user"
	"identity-api/Infrastructure/databases/postgres/entities"
	base "identity-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

func ProvideUserPersistenceMapper() base.PersistenceMapper[entities.UserEntity, models.UserModel] {
	return NewUserPersistenceMapper()
}

var WireSet = wire.NewSet(
	ProvideUserPersistenceMapper,
)
