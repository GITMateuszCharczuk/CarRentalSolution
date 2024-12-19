package repository_interfaces

import (
	"context"
	models "rental-api/Domain/models/domestic"
)

type CarOfferCommandRepository interface {
	CreateCarOffer(ctx context.Context, carOffer *models.CarOfferModel, tags []string, imageIds []string) (*string, error)
	UpdateCarOffer(ctx context.Context, carOffer *models.CarOfferModel, tags []string, imageIds []string) error
	DeleteCarOffer(ctx context.Context, id string) error
}
