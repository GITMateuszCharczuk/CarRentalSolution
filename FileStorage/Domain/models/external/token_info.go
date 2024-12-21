package models

import "file-storage/Domain/constants"

type TokenInfo struct {
	Valid bool                `json:"valid"`
	Roles []constants.JWTRole `json:"roles"`
}
