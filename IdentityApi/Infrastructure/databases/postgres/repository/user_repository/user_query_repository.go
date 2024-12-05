package repository

import (
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/user"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"
)

type UserQueryRepositoryImpl struct {
	*base.QueryRepository[entities.UserEntity, string, models.UserModel]
	mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]
}

func NewUserQueryRepositoryImpl(postgresDatabase *postgres_db.PostgresDatabase, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) *UserQueryRepositoryImpl {
	return &UserQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.UserEntity, string, models.UserModel](postgresDatabase.DB, mapper),
		mapper:          mapper,
	}
}

func (r *UserQueryRepositoryImpl) GetUserByID(id string) (*models.UserModel, error) {
	userModel, err := r.GetById(id)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (r *UserQueryRepositoryImpl) GetUserByEmail(email string) (*models.UserModel, error) {
	userModel, err := r.GetFirstByProp("email", email)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (r *UserQueryRepositoryImpl) GetUsersByRoles(roles ...constants.JWTRole) ([]*models.UserModel, error) {
	roleEntities := make([]entities.JWTRoleEntity, len(roles))
	for i, role := range roles {
		roleEntities[i] = entities.JWTRoleEntity(role)
	}
	return r.GetAllByPropValues("roles", roleEntities)
}
