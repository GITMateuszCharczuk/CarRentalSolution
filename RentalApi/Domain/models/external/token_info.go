package models

import "rental-api/Domain/constants"

type TokenInfo struct {
	Valid bool                `json:"valid"`
	Roles []constants.JWTRole `json:"roles"`
}
