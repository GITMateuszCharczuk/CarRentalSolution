package contract

import (
	models "rental-api/Domain/models/domestic"
	responses "rental-api/Domain/responses"
)

type GetAllImagesResponse struct {
	Images []models.CarOfferImageModel `json:"images" example:"https://example.com/image1.jpg,https://example.com/image2.jpg" swaggertype:"array,string"`
	responses.BaseResponse
}
