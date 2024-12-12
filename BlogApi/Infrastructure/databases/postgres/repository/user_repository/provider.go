package repository

import (
	models "identity-api/Domain/models/user"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"

	"github.com/google/wire"
)

func ProvideUserQueryRepository(postgresDatabase *postgres_db.PostgresDatabase, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) repository_interfaces.UserQueryRepository {
	return NewUserQueryRepositoryImpl(postgresDatabase, mapper)
}

func ProvideUserCommandRepository(postgresDatabase *postgres_db.PostgresDatabase, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) repository_interfaces.UserCommandRepository {
	return NewUserCommandRepositoryImpl(postgresDatabase, mapper)
}

var WireSet = wire.NewSet(ProvideUserQueryRepository, ProvideUserCommandRepository)
