package models

import "identity-api/Domain/constants"

type TokenInfo struct {
	Valid bool                `json:"valid"`
	Roles []constants.JWTRole `json:"roles"`
}
