package mappers

import (
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/user"
	"identity-api/Infrastructure/databases/postgres/entities"

	"github.com/google/uuid"
)

type UserMapper struct{}

func NewUserPersistenceMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) MapToModel(entity entities.UserEntity) models.UserModel {
	return models.UserModel{
		ID:           entity.ID.String(),
		Roles:        convertRoles(entity.Roles),
		Name:         entity.Name,
		Surname:      entity.Surname,
		PhoneNumber:  entity.PhoneNumber,
		EmailAddress: entity.EmailAddress,
		Password:     entity.Password,
		Address:      entity.Address,
		PostalCode:   entity.PostalCode,
		City:         entity.City,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

func (m *UserMapper) MapToEntity(model models.UserModel) entities.UserEntity {
	return entities.UserEntity{
		ID:           uuid.New(),
		Roles:        convertRolesBack(model.Roles),
		Name:         model.Name,
		Surname:      model.Surname,
		PhoneNumber:  model.PhoneNumber,
		EmailAddress: model.EmailAddress,
		Password:     model.Password,
		Address:      model.Address,
		PostalCode:   model.PostalCode,
		City:         model.City,
	}
}

func convertRoles(roles []entities.JWTRoleEntity) []constants.JWTRole {
	jwtRoles := make([]constants.JWTRole, len(roles))
	for i, role := range roles {
		jwtRoles[i] = constants.JWTRole(role)
	}
	return jwtRoles
}

func convertRolesBack(roles []constants.JWTRole) []entities.JWTRoleEntity {
	roleEntities := make([]entities.JWTRoleEntity, len(roles))
	for i, role := range roles {
		roleEntities[i] = entities.JWTRoleEntity(role)
	}
	return roleEntities
}
