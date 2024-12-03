package repository

import (
	models "identity-api/Domain/models/user"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"

	"gorm.io/gorm"
)

type UserCommandRepositoryImpl struct {
	*base.CommandRepository[entities.UserEntity, string, models.UserModel]
}

func NewUserCommandRepositoryImpl(dbContext *gorm.DB, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) *UserCommandRepositoryImpl {
	return &UserCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.UserEntity, string, models.UserModel](dbContext, mapper),
	}
}

func (r *UserCommandRepositoryImpl) CreateUser(user *models.UserModel) error {
	_, err := r.Add(*user)
	return err
}

func (r *UserCommandRepositoryImpl) UpdateUser(user *models.UserModel) error {
	_, err := r.Update(*user)
	return err
}

func (r *UserCommandRepositoryImpl) DeleteUser(id string) error {
	return r.Delete(id)
}
