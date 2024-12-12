package contract

import (
	models "identity-api/Domain/models/domestic"
	ex
	"identity-api/Domain/pagination"
	responses "identity-api/Domain/responses"
)

type GetAllUsersResponse struct {
	responses.BaseResponse
	pagination.PaginatedResult[models.UserSecureInfo] `json:",inline"`
}

type GetAllUsersResponse200 struct {
	StatusCode int    `json:"status_code" example:"200"`
	Message    string `json:"message" example:"Users retrieved successfully"`
	Data       struct {
		Items []struct {
			ID           string   `json:"id" example:"12345"`
			Name         string   `json:"name" example:"John"`
			Surname      string   `json:"surname" example:"Doe"`
			PhoneNumber  string   `json:"phone_number" example:"+1234567890"`
			EmailAddress string   `json:"email_address" example:"user@example.com"`
			Address      string   `json:"address" example:"123 Main St"`
			PostalCode   string   `json:"postal_code" example:"12345"`
			City         string   `json:"city" example:"New York"`
			Roles        []string `json:"roles" example:"user,admin"`
		} `json:"items"`
		TotalItems  int `json:"total_items" example:"100"`
		CurrentPage int `json:"current_page" example:"1"`
		PageSize    int `json:"page_size" example:"10"`
		TotalPages  int `json:"total_pages" example:"10"`
	} `json:"data"`
}

type GetAllUsersResponse400 struct {
	Success bool               `json:"success" example:"false" swaggertype:"boolean"`
	Message string             `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Data    []dummyArrayObject `json:"data" swaggertype:"array,object"`
}

type GetAllUsersResponse401 struct {
	Success bool               `json:"success" example:"false" swaggertype:"boolean"`
	Message string             `json:"message" example:"Invalid or expired token" swaggertype:"string"`
	Data    []dummyArrayObject `json:"data" swaggertype:"array,object"`
}

type GetAllUsersResponse403 struct {
	Success bool               `json:"success" example:"false" swaggertype:"boolean"`
	Message string             `json:"message" example:"Insufficient privileges" swaggertype:"string"`
	Data    []dummyArrayObject `json:"data" swaggertype:"array,object"`
}

type GetAllUsersResponse500 struct {
	Success bool               `json:"success" example:"false" swaggertype:"boolean"`
	Message string             `json:"message" example:"Internal server error during retrieval" swaggertype:"string"`
	Data    []dummyArrayObject `json:"data" swaggertype:"array,object"`
}

type dummyArrayObject struct{}
