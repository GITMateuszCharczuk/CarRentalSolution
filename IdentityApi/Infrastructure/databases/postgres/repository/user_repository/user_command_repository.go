package repository

import (
	models "identity-api/Domain/models/user"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
)

type UserCommandRepositoryImpl struct {
	*base.CommandRepository[entities.UserEntity, string, models.UserModel]
}

func NewUserCommandRepositoryImpl(postgresDatabase *postgres_db.PostgresDatabase, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) *UserCommandRepositoryImpl {
	return &UserCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.UserEntity, string, models.UserModel](postgresDatabase.DB, mapper),
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
