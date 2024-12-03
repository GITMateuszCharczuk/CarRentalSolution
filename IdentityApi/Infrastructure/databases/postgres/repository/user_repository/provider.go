package repository

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	models "identity-api/Domain/models/user"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
)

func ProvideUserQueryRepository(dbContext *gorm.DB, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) repository_interfaces.UserQueryRepository {
	return NewUserQueryRepositoryImpl(dbContext, mapper)
}

func ProvideUserCommandRepository(dbContext *gorm.DB, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) repository_interfaces.UserCommandRepository {
	return NewUserCommandRepositoryImpl(dbContext, mapper)
}

var WireSet = wire.NewSet(ProvideUserQueryRepository, ProvideUserCommandRepository)
