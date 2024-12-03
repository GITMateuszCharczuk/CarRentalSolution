package entities

import (
	"time"

	"github.com/google/uuid"
)

type JWTRoleEntity string

const (
	User       JWTRoleEntity = "user"
	Admin      JWTRoleEntity = "admin"
	SuperAdmin JWTRoleEntity = "superadmin"
)

type UserEntity struct {
	ID           uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	Roles        []JWTRoleEntity `gorm:"type:user_role[]" json:"roles"`
	Name         string          `json:"name"`
	Surname      string          `json:"surname"`
	PhoneNumber  string          `json:"phone_number"`
	EmailAddress string          `gorm:"unique" json:"email_address"`
	Password     string          `json:"password"`
	Address      string          `json:"address"`
	PostalCode   string          `json:"postal_code"`
	City         string          `json:"city"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}
