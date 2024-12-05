package contract

import (
	models "identity-api/Domain/models/user"
	responses "identity-api/Domain/responses"
)

type GetAllUsersResponse struct {
	responses.BaseResponse
	Users []models.UserSecureInfo `json:"users" swaggertype:"array,object"`
}
