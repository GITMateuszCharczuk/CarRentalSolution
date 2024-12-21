package models

import "email-service/Domain/constants"

type TokenInfo struct {
	Valid bool                `json:"valid"`
	Roles []constants.JWTRole `json:"roles"`
}
