package mappers

import (
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/user"
	"identity-api/Infrastructure/databases/postgres/entities"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserMapper struct{}

func NewUserPersistenceMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) MapToModel(entity entities.UserEntity) models.UserModel {
	log.Println(entity)
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
	var id uuid.UUID
	if model.ID == "" {
		id = uuid.New()
	} else {
		id, _ = uuid.Parse(model.ID)
	}

	return entities.UserEntity{
		ID:           id,
		Roles:        convertRolesBack(model.Roles),
		Name:         model.Name,
		Surname:      model.Surname,
		PhoneNumber:  model.PhoneNumber,
		EmailAddress: model.EmailAddress,
		Password:     model.Password,
		Address:      model.Address,
		PostalCode:   model.PostalCode,
		City:         model.City,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func convertRoles(roles entities.RoleArray) []constants.JWTRole {
	jwtRoles := make([]constants.JWTRole, len(roles))
	for i, role := range roles {
		jwtRoles[i] = constants.JWTRole(role)
	}
	return jwtRoles
}

func convertRolesBack(roles []constants.JWTRole) entities.RoleArray {
	if len(roles) == 0 {
		return entities.RoleArray{entities.User}
	}

	roleEntities := make([]entities.JWTRoleEntity, len(roles))
	for i, role := range roles {
		roleEntities[i] = entities.JWTRoleEntity(role)
	}
	return entities.RoleArray(roleEntities)
}
