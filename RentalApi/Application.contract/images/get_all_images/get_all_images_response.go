package contract

import (
	models "rental-api/Domain/models/domestic"
	responses "rental-api/Domain/responses"
)

type GetAllImagesResponse struct {
	Items []models.CarOfferImageModel `json:"items" example:"https://example.com/image1.jpg,https://example.com/image2.jpg" swaggertype:"array,string"`
	responses.BaseResponse
}

type GetAllImagesResponse200 struct {
	Items   []models.CarOfferImageModel `json:"items" example:"https://example.com/image1.jpg,https://example.com/image2.jpg" swaggertype:"array,string"`
	Success bool                        `json:"success" example:"true" swaggertype:"boolean"`
	Message string                      `json:"message" example:"Images retrieved successfully" swaggertype:"string"`
}

type GetAllImagesResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type GetAllImagesResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type GetAllImagesResponse403 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Forbidden - Not authorized" swaggertype:"string"`
}

type GetAllImagesResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer not found" swaggertype:"string"`
}

type GetAllImagesResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Server error during retrieval" swaggertype:"string"`
}
