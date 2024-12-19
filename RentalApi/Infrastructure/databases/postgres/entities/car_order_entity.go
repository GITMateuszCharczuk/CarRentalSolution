package entities

import (
	"time"

	"rental-api/Domain/constants"

	"github.com/google/uuid"
)

type CarOrderEntity struct {
	ID               uuid.UUID                `gorm:"type:uuid;primaryKey" json:"id"`
	UserId           uuid.UUID                `gorm:"type:uuid;index" json:"user_id"`
	CarOfferId       *uuid.UUID               `gorm:"type:uuid;index" json:"car_offer_id"`
	StartDate        time.Time                `json:"start_date"`
	EndDate          time.Time                `json:"end_date"`
	NumOfDrivers     int                      `json:"num_of_drivers"`
	DeliveryLocation string                   `json:"delivery_location"`
	ReturnLocation   string                   `json:"return_location"`
	TotalCost        float64                  `json:"total_cost"`
	CreatedAt        time.Time                `json:"created_at"`
	UpdatedAt        time.Time                `json:"updated_at"`
	Status           constants.CarOrderStatus `gorm:"type:varchar(20);default:'PENDING'" json:"status"`

	// Navigation Properties
	CarOffer CarOfferEntity `gorm:"foreignKey:CarOfferId" json:"car_offer"`
}
