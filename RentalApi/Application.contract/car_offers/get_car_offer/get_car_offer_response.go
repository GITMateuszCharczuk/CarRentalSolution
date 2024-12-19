package contract

import (
	models "rental-api/Domain/models/domestic"
	responses "rental-api/Domain/responses"
)

type GetCarOfferResponse struct {
	responses.BaseResponse
	CarOffer models.CarOfferModel `json:"car_offer" swaggertype:"object"`
}

type GetCarOfferResponse200 struct {
	Success  bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message  string `json:"message" example:"Car offer retrieved successfully" swaggertype:"string"`
	CarOffer struct {
		Id                 string   `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		Heading            string   `json:"heading" example:"Car Offer Title" swaggertype:"string"`
		ShortDescription   string   `json:"shortDescription" example:"Short description" swaggertype:"string"`
		FeaturedImageUrl   string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
		UrlHandle          string   `json:"urlHandle" example:"car-offer-title" swaggertype:"string"`
		Horsepower         string   `json:"horsepower" example:"300" swaggertype:"string"`
		YearOfProduction   int      `json:"yearOfProduction" example:"2023" swaggertype:"integer"`
		EngineDetails      string   `json:"engineDetails" example:"2.0L Turbo" swaggertype:"string"`
		DriveDetails       string   `json:"driveDetails" example:"AWD" swaggertype:"string"`
		GearboxDetails     string   `json:"gearboxDetails" example:"Automatic" swaggertype:"string"`
		DeliveryLocation   string   `json:"deliveryLocation" example:"Airport" swaggertype:"string"`
		Visible            bool     `json:"visible" example:"true" swaggertype:"boolean"`
		OneNormalDayPrice  float64  `json:"oneNormalDayPrice" example:"100.00" swaggertype:"number"`
		OneWeekendDayPrice float64  `json:"oneWeekendDayPrice" example:"150.00" swaggertype:"number"`
		OneWeekPrice       float64  `json:"oneWeekPrice" example:"600.00" swaggertype:"number"`
		OneMonthPrice      float64  `json:"oneMonthPrice" example:"2000.00" swaggertype:"number"`
		Tags               []string `json:"tags" example:"[\"luxury\",\"sports\"]" swaggertype:"array,string"`
		ImageUrls          []string `json:"imageUrls" example:"[\"https://example.com/image1.jpg\"]" swaggertype:"array,string"`
	} `json:"car_offer"`
}

type GetCarOfferResponse400 struct {
	Success  bool                 `json:"success" example:"false" swaggertype:"boolean"`
	Message  string               `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	CarOffer models.CarOfferModel `json:"car_offer" swaggertype:"object"`
}

type GetCarOfferResponse404 struct {
	Success  bool                 `json:"success" example:"false" swaggertype:"boolean"`
	Message  string               `json:"message" example:"Car offer not found" swaggertype:"string"`
	CarOffer models.CarOfferModel `json:"car_offer" swaggertype:"object"`
}

type GetCarOfferResponse500 struct {
	Success  bool                 `json:"success" example:"false" swaggertype:"boolean"`
	Message  string               `json:"message" example:"Internal server error while retrieving car offer" swaggertype:"string"`
	CarOffer models.CarOfferModel `json:"car_offer" swaggertype:"object"`
}
