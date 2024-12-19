package repository_interfaces

import (
	"context"
	models "rental-api/Domain/models/domestic"
	"rental-api/Infrastructure/databases/postgres/entities"
)

type CarTagCommandRepository interface {
	AddTagToCarOffer(ctx context.Context, tag *models.CarOfferTagModel, carOfferEntity entities.CarOfferEntity) (*models.CarOfferTagModel, error)
	AddTagsToCarOffer(ctx context.Context, carOfferId string, tagNames []string) error
	ModifyTagsForCarOffer(ctx context.Context, carOfferId string, newTagNames []string) error
	CleanupUnusedTags(ctx context.Context) error
}
