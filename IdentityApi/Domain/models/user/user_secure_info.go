package models

type UserSecureInfo struct {
	UserInfo
	ID    string   `json:"id" example:"12345" swaggertype:"string"`
	Roles []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}
