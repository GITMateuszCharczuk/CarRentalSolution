package models

type TokenInfo struct {
	Valid bool     `json:"valid"`
	Roles []string `json:"roles"`
}
