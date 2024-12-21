package models

type CarOfferModel struct {
	Id                 string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Heading            string  `json:"heading" example:"Car Offer Title" swaggertype:"string"`
	ShortDescription   string  `json:"short_description" example:"Short description" swaggertype:"string"`
	UrlHandle          string  `json:"url_handle" example:"car-offer-title" swaggertype:"string"`
	Horsepower         string  `json:"horsepower" example:"300" swaggertype:"string"`
	YearOfProduction   int     `json:"year_of_production" example:"2023" swaggertype:"integer"`
	EngineDetails      string  `json:"engine_details" example:"2.0L Turbo" swaggertype:"string"`
	DriveDetails       string  `json:"drive_details" example:"AWD" swaggertype:"string"`
	GearboxDetails     string  `json:"gearbox_details" example:"Automatic" swaggertype:"string"`
	PublishedDate      string  `json:"published_date" example:"2023-12-12" swaggertype:"string"`
	Visible            bool    `json:"visible" example:"true" swaggertype:"boolean"`
	OneNormalDayPrice  float64 `json:"one_normal_day_price" example:"100.00" swaggertype:"number"`
	OneWeekendDayPrice float64 `json:"one_weekend_day_price" example:"150.00" swaggertype:"number"`
	OneWeekPrice       float64 `json:"one_week_price" example:"600.00" swaggertype:"number"`
	OneMonthPrice      float64 `json:"one_month_price" example:"2000.00" swaggertype:"number"`
	CustodianId        string  `json:"custodian_id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	CustodianEmail     string  `json:"custodian_email" example:"custodian@example.com" swaggertype:"string"`
	CreatedAt          string  `json:"created_at" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	UpdatedAt          string  `json:"updated_at" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	FeaturedImageUrl   string  `json:"featured_image_url" example:"https://example.com/image.jpg" swaggertype:"string"`
}
