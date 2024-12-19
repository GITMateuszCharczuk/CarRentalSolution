package entities

import (
	"time"

	"github.com/google/uuid"
)

type CarImageEntity struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ImageId    string    `json:"image_id"`
	CarOfferId uuid.UUID `gorm:"type:uuid;index" json:"car_offer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Navigation Properties
	CarOffer CarOfferEntity `gorm:"foreignKey:CarOfferId" json:"car_offer"`
}
