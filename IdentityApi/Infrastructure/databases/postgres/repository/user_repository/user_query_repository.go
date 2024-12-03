package repository

import (
	models "identity-api/Domain/models/user"
	"identity-api/Infrastructure/databases/postgres/entities"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	base "identity-api/Infrastructure/databases/postgres/repository/base"

	"gorm.io/gorm"
)

type UserQueryRepositoryImpl struct {
	*base.QueryRepository[entities.UserEntity, string, models.UserModel]
}

func NewUserQueryRepositoryImpl(dbContext *gorm.DB, mapper mappers.PersistenceMapper[entities.UserEntity, models.UserModel]) *UserQueryRepositoryImpl {
	return &UserQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.UserEntity, string, models.UserModel](dbContext, mapper),
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
	userModel, err := r.GetByProp("email", email)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}
