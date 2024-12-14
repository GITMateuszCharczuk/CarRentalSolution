package services

import "identity-api/Domain/constants"

func IsAdminOrSuperAdmin(roles []constants.JWTRole) bool {
	for _, role := range roles {
		if role == constants.Admin || role == constants.SuperAdmin {
			return true
		}
	}
	return false
}

func IsRole(role constants.JWTRole, roles []constants.JWTRole) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
