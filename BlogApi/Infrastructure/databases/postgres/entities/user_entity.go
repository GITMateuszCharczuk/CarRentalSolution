package entities

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type JWTRoleEntity string

const (
	User       JWTRoleEntity = "user"
	Admin      JWTRoleEntity = "admin"
	SuperAdmin JWTRoleEntity = "superadmin"
)

type RoleArray []JWTRoleEntity

func (r *RoleArray) Scan(value interface{}) error {
	if value == nil {
		*r = []JWTRoleEntity{User}
		return nil
	}

	var arr pq.StringArray
	err := arr.Scan(value)
	if err != nil {
		return err
	}

	*r = make([]JWTRoleEntity, len(arr))
	for i, v := range arr {
		(*r)[i] = JWTRoleEntity(v)
	}
	return nil
}

func (r RoleArray) Value() (driver.Value, error) {
	if len(r) == 0 {
		return pq.StringArray{string(User)}.Value()
	}

	strArr := make([]string, len(r))
	for i, v := range r {
		strArr[i] = string(v)
	}
	return pq.StringArray(strArr).Value()
}

type UserEntity struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Roles        RoleArray `gorm:"type:text[]" json:"roles"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	PhoneNumber  string    `json:"phone_number"`
	EmailAddress string    `gorm:"unique" json:"email_address"`
	Password     string    `json:"password"`
	Address      string    `json:"address"`
	PostalCode   string    `json:"postal_code"`
	City         string    `json:"city"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
