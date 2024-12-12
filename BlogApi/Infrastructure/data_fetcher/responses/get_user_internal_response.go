package datafetcher

import models "identity-api/Domain/models/external"

type GetUserInternalResponse struct {
	Success  bool            `json:"success"`
	Message  string          `json:"message"`
	UserInfo models.UserInfo `json:"user_info"`
}
