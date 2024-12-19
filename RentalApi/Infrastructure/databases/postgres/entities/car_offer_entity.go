package entities

import (
	"time"

	"github.com/google/uuid"
)

type CarOfferEntity struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Heading            string    `json:"heading"`
	ShortDescription   string    `json:"short_description"`
	FeaturedImageUrl   string    `json:"featured_image_url"`
	UrlHandle          string    `json:"url_handle"`
	Horsepower         string    `json:"horsepower"`
	YearOfProduction   int       `json:"year_of_production"`
	EngineDetails      string    `json:"engine_details"`
	DriveDetails       string    `json:"drive_details"`
	GearboxDetails     string    `json:"gearbox_details"`
	PublishedDate      time.Time `json:"published_date"`
	Visible            bool      `json:"visible"`
	OneNormalDayPrice  float64   `json:"one_normal_day_price"`
	OneWeekendDayPrice float64   `json:"one_weekend_day_price"`
	OneWeekPrice       float64   `json:"one_week_price"`
	OneMonthPrice      float64   `json:"one_month_price"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	CustodianId        uuid.UUID `json:"custodian_id"`
	CustodianEmail     string    `json:"custodian_email"`

	// Navigation Properties
	Tags   []CarTagEntity   `gorm:"many2many:car_offer_tags;foreignKey:ID;joinForeignKey:CarOfferID;References:ID;joinReferences:TagID" json:"tags"`
	Images []CarImageEntity `gorm:"foreignKey:CarOfferID;constraint:OnDelete:CASCADE;" json:"images"`
	Orders []CarOrderEntity `gorm:"foreignKey:CarOfferID;constraint:OnDelete:SET NULL;" json:"orders"`
}
