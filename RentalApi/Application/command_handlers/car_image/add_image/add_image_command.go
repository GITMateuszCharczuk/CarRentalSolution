package commands

import (
	models "rental-api/Domain/models/external"
)

type AddImageCommand struct {
	CarOfferId      string `json:"carOfferId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	ImageId         string `json:"image_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	models.JwtToken `json:",inline"`
}
