package contract

import (
	models "identity-api/Domain/models/user"
	"identity-api/Domain/responses"
)

type GetUserInfoResponse struct {
	responses.BaseResponse
	models.UserSecureInfo
}
