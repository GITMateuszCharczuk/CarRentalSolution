package models

import "time"

type UserModel struct {
	ID        string
	Email     string
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time
}
