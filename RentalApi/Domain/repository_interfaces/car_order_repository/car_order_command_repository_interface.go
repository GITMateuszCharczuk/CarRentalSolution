package repository_interfaces

import (
	"context"
	models "rental-api/Domain/models/domestic"
)

type CarOrderCommandRepository interface {
	CreateCarOrder(ctx context.Context, carOrder *models.CarOrderModel) (*string, error)
	UpdateCarOrder(ctx context.Context, carOrder *models.CarOrderModel) error
	DeleteCarOrder(ctx context.Context, id string) error
}
