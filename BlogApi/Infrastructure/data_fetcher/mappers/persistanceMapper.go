package datafetcher

import (
	models "identity-api/Domain/models/external"
	responses "identity-api/Infrastructure/data_fetcher/responses"
)

func MapToTokenInfo(tokenInfo responses.ValidateTokenResponse) models.TokenInfo {
	return models.TokenInfo{
		Valid: tokenInfo.Valid,
		Roles: tokenInfo.Roles,
	}
}

func MapToUserInfo(userInfo responses.GetUserInternalResponse) models.UserInfo {
	return userInfo.UserInfo
}
