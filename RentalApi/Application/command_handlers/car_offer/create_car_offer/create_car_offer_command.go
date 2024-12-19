package commands

import (
	models "rental-api/Domain/models/external"
)

type CreateCarOfferCommand struct {
	Heading            string   `json:"heading" binding:"required" example:"Car Offer Title" swaggertype:"string" validate:"required"`
	ShortDescription   string   `json:"shortDescription" binding:"required" example:"Short description" swaggertype:"string" validate:"required"`
	FeaturedImageUrl   string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle          string   `json:"urlHandle" binding:"required" example:"car-offer-title" swaggertype:"string" validate:"required"`
	Horsepower         string   `json:"horsepower" example:"300" swaggertype:"string"`
	YearOfProduction   int      `json:"yearOfProduction" example:"2023" swaggertype:"integer"`
	EngineDetails      string   `json:"engineDetails" example:"2.0L Turbo" swaggertype:"string"`
	DriveDetails       string   `json:"driveDetails" example:"AWD" swaggertype:"string"`
	GearboxDetails     string   `json:"gearboxDetails" example:"Automatic" swaggertype:"string"`
	Visible            bool     `json:"visible" example:"true" swaggertype:"boolean"`
	OneNormalDayPrice  float64  `json:"oneNormalDayPrice" example:"100.00" swaggertype:"number"`
	OneWeekendDayPrice float64  `json:"oneWeekendDayPrice" example:"150.00" swaggertype:"number"`
	OneWeekPrice       float64  `json:"oneWeekPrice" example:"600.00" swaggertype:"number"`
	OneMonthPrice      float64  `json:"oneMonthPrice" example:"2000.00" swaggertype:"number"`
	Tags               []string `json:"tags" example:"[\"luxury\",\"sports\"]" swaggertype:"array,string"`
	ImageUrls          []string `json:"imageUrls" example:"[\"https://example.com/image1.jpg\"]" swaggertype:"array,string"`
	models.JwtToken    `json:",inline"`
}
