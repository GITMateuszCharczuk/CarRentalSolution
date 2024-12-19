package repository_interfaces

import (
	"context"
)

type CarImageCommandRepository interface {
	AddImageToCarOffer(ctx context.Context, carOfferId string, imageId string) (*string, error)
	AddImagesToCarOffer(ctx context.Context, carOfferId string, imageIds []string) error
	ModifyImagesForCarOffer(ctx context.Context, carOfferId string, imageIds []string) error
	DeleteImageFromCarOffer(ctx context.Context, carOfferId string, imageId string) error
}
