package contract

import models "rental-api/Domain/models/external"

type UpdateCarOfferRequest struct {
	Id                 string   `json:"-" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	Heading            string   `json:"heading" binding:"required" example:"Updated Car Offer Title" swaggertype:"string" validate:"required"`
	ShortDescription   string   `json:"shortDescription" binding:"required" example:"Updated short description" swaggertype:"string" validate:"required"`
	FeaturedImageUrl   string   `json:"featuredImageUrl" example:"https://example.com/updated-image.jpg" swaggertype:"string"`
	UrlHandle          string   `json:"urlHandle" binding:"required" example:"updated-car-offer-title" swaggertype:"string" validate:"required"`
	Horsepower         string   `json:"horsepower" example:"350" swaggertype:"string"`
	YearOfProduction   int      `json:"yearOfProduction" example:"2023" swaggertype:"integer"`
	EngineDetails      string   `json:"engineDetails" example:"3.0L Twin-Turbo" swaggertype:"string"`
	DriveDetails       string   `json:"driveDetails" example:"RWD" swaggertype:"string"`
	GearboxDetails     string   `json:"gearboxDetails" example:"8-Speed Automatic" swaggertype:"string"`
	PublishedDate      string   `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
	Visible            bool     `json:"visible" example:"true" swaggertype:"boolean"`
	OneNormalDayPrice  float64  `json:"oneNormalDayPrice" example:"120.00" swaggertype:"number"`
	OneWeekendDayPrice float64  `json:"oneWeekendDayPrice" example:"180.00" swaggertype:"number"`
	OneWeekPrice       float64  `json:"oneWeekPrice" example:"700.00" swaggertype:"number"`
	OneMonthPrice      float64  `json:"oneMonthPrice" example:"2500.00" swaggertype:"number"`
	Tags               []string `json:"tags" example:"[\"luxury\",\"sports\"]" swaggertype:"array,string"`
	models.JwtToken    `json:",inline"`
}
