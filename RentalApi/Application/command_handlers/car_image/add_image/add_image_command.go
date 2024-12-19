package commands

import (
	models "rental-api/Domain/models/external"
)

type AddImageCommand struct {
	CarOfferId      string `json:"carOfferId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	Url             string `json:"url" binding:"required" example:"https://example.com/image1.jpg" swaggertype:"string" validate:"required"`
	IsFeatured      bool   `json:"isFeatured" example:"false" swaggertype:"boolean"`
	Order           int    `json:"order" example:"1" swaggertype:"integer"`
	models.JwtToken `json:",inline"`
}
