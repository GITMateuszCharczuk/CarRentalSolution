package datafetcher

import "identity-api/Domain/constants"

func ConvertRolesToJWTRole(roles []string) []constants.JWTRole {
	roleJWTRoles := make([]constants.JWTRole, len(roles))
	for i, role := range roles {
		roleJWTRoles[i] = constants.JWTRole(role)
	}
	return roleJWTRoles
}
