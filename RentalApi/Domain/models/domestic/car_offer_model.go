package models

type CarOfferModel struct {
	Id                 string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Heading            string  `json:"heading" example:"Car Offer Title" swaggertype:"string"`
	ShortDescription   string  `json:"shortDescription" example:"Short description" swaggertype:"string"`
	UrlHandle          string  `json:"urlHandle" example:"car-offer-title" swaggertype:"string"`
	Horsepower         string  `json:"horsepower" example:"300" swaggertype:"string"`
	YearOfProduction   int     `json:"yearOfProduction" example:"2023" swaggertype:"integer"`
	EngineDetails      string  `json:"engineDetails" example:"2.0L Turbo" swaggertype:"string"`
	DriveDetails       string  `json:"driveDetails" example:"AWD" swaggertype:"string"`
	GearboxDetails     string  `json:"gearboxDetails" example:"Automatic" swaggertype:"string"`
	PublishedDate      string  `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
	Visible            bool    `json:"visible" example:"true" swaggertype:"boolean"`
	OneNormalDayPrice  float64 `json:"oneNormalDayPrice" example:"100.00" swaggertype:"number"`
	OneWeekendDayPrice float64 `json:"oneWeekendDayPrice" example:"150.00" swaggertype:"number"`
	OneWeekPrice       float64 `json:"oneWeekPrice" example:"600.00" swaggertype:"number"`
	OneMonthPrice      float64 `json:"oneMonthPrice" example:"2000.00" swaggertype:"number"`
	CustodianId        string  `json:"custodianId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	CustodianEmail     string  `json:"custodianEmail" example:"custodian@example.com" swaggertype:"string"`
	CreatedAt          string  `json:"createdAt" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	UpdatedAt          string  `json:"updatedAt" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	FeaturedImageUrl   string  `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
}
