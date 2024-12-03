package repository_interfaces

import models "identity-api/Domain/models/user"

type UserQueryRepository interface {
	GetUserByID(id string) (*models.UserModel, error)
	GetUserByEmail(email string) (*models.UserModel, error)
}
