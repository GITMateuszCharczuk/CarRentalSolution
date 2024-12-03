package services

import "identity-api/Domain/constants"

func ConvertRolesToString(roles []constants.JWTRole) []string {
	roleStrings := make([]string, len(roles))
	for i, role := range roles {
		roleStrings[i] = string(role)
	}
	return roleStrings
}

func ConvertRolesToJWTRole(roles []string) []constants.JWTRole {
	roleJWTRoles := make([]constants.JWTRole, len(roles))
	for i, role := range roles {
		roleJWTRoles[i] = constants.JWTRole(role)
	}
	return roleJWTRoles
}

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
