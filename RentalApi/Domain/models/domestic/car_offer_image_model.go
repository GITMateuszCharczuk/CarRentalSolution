package models

type CarOfferImageModel struct { //TODO change name
	Id         string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	ImageId    string `json:"url" example:"https://example.com/image1.jpg" swaggertype:"string"`
	CarOfferId string `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	CreatedAt  string `json:"createdAt" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	UpdatedAt  string `json:"updatedAt" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
}
