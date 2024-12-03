package models

import (
	"identity-api/Domain/constants"
	"time"
)

type UserModel struct {
	ID           string              `json:"id"`
	Roles        []constants.JWTRole `json:"roles"`
	Name         string              `json:"name"`
	Surname      string              `json:"surname"`
	PhoneNumber  string              `json:"phone_number"`
	EmailAddress string              `json:"email_address"`
	Password     string              `json:"password"`
	Address      string              `json:"address"`
	PostalCode   string              `json:"postal_code"`
	City         string              `json:"city"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}
