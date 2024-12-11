package repository_interfaces

import (
	"identity-api/Domain/constants"
	models "identity-api/Domain/models/user"
	"identity-api/Domain/pagination"
	"identity-api/Domain/sorting"
)

type UserQueryRepository interface {
	GetUserByID(id string) (*models.UserModel, error)
	GetUserByEmail(email string) (*models.UserModel, error)
	GetUsersByRoles(roles []constants.JWTRole, pagination *pagination.Pagination, sorting *sorting.Sortable) (*pagination.PaginatedResult[models.UserModel], error)
}
