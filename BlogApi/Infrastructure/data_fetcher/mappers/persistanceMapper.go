package datafetcher

import (
	models "identity-api/Domain/models/external"
	responses "identity-api/Infrastructure/data_fetcher/responses"
	utils "identity-api/Infrastructure/data_fetcher/utils"
)

func MapToTokenInfo(tokenInfo responses.ValidateTokenResponse) models.TokenInfo {
	return models.TokenInfo{
		Valid: tokenInfo.Valid,
		Roles: utils.ConvertRolesToJWTRole(tokenInfo.Roles),
	}
}

func MapToUserInfo(userInfo responses.GetUserInternalResponse) models.UserInfo {
	return models.UserInfo{
		Name:         userInfo.UserResponseInfo.Name,
		ID:           userInfo.UserResponseInfo.ID,
		Surname:      userInfo.UserResponseInfo.Surname,
		PhoneNumber:  userInfo.UserResponseInfo.PhoneNumber,
		EmailAddress: userInfo.UserResponseInfo.EmailAddress,
		Address:      userInfo.UserResponseInfo.Address,
		PostalCode:   userInfo.UserResponseInfo.PostalCode,
		City:         userInfo.UserResponseInfo.City,
		Roles:        utils.ConvertRolesToJWTRole(userInfo.UserResponseInfo.Roles),
	}
}
