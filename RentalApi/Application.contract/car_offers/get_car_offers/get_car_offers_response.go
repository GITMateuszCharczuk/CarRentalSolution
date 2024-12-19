package contract

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Domain/pagination"
	responses "rental-api/Domain/responses"
)

type GetCarOffersResponse struct {
	responses.BaseResponse
	pagination.PaginatedResult[models.CarOfferModel] `json:",inline"`
}

type GetCarOffersResponse200 struct {
	Success     bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message     string `json:"message" example:"Car offers retrieved successfully" swaggertype:"string"`
	TotalItems  int    `json:"total_items" example:"100"`
	CurrentPage int    `json:"current_page" example:"1"`
	PageSize    int    `json:"page_size" example:"10"`
	TotalPages  int    `json:"total_pages" example:"10"`
	Items       []struct {
		Id               string   `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		Heading          string   `json:"heading" example:"Car Offer Title" swaggertype:"string"`
		ShortDescription string   `json:"shortDescription" example:"Short description" swaggertype:"string"`
		FeaturedImageUrl string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
		UrlHandle        string   `json:"urlHandle" example:"car-offer-title" swaggertype:"string"`
		PublishedDate    string   `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
		Visible          bool     `json:"visible" example:"true" swaggertype:"boolean"`
		Tags             []string `json:"tags" example:"[\"luxury\",\"sports\"]" swaggertype:"array,string"`
	} `json:"items"`
}

type GetCarOffersResponse400 struct {
	Success bool                   `json:"success" example:"false" swaggertype:"boolean"`
	Message string                 `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Items   []models.CarOfferModel `json:"items" swaggertype:"array,object"`
}

type GetCarOffersResponse500 struct {
	Success bool                   `json:"success" example:"false" swaggertype:"boolean"`
	Message string                 `json:"message" example:"Internal server error while retrieving car offers" swaggertype:"string"`
	Items   []models.CarOfferModel `json:"items" swaggertype:"array,object"`
}
