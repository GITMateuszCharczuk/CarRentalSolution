package repository_interfaces

import models "identity-api/Domain/models/external"

type UserCommandRepository interface {
	CreateUser(user *models.UserModel) error
	UpdateUser(user *models.UserModel) error
	DeleteUser(id string) error
}
