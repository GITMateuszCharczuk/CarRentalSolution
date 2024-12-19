package entities

import (
	"time"

	"github.com/google/uuid"
)

type CarTagEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"unique" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Navigation Properties
	CarOffers []CarOfferEntity `gorm:"many2many:car_offer_tags;foreignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:CarOfferID" json:"car_offers"`
}
