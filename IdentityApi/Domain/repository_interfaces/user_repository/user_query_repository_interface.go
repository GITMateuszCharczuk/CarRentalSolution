package repository_interfaces

import (
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/user"
)

type UserQueryRepository interface {
	GetUserByID(id string) (*models.UserModel, error)
	GetUserByEmail(email string) (*models.UserModel, error)
	GetUsersByRoles(roles ...constants.JWTRole) ([]*models.UserModel, error)
}
