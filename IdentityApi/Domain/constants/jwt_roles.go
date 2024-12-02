package constants

type JWTRole string

const ( //TODO pozmieniac w innych miejscach
	User       JWTRole = "user"
	Admin      JWTRole = "admin"
	SuperAdmin JWTRole = "superadmin"
)

func (r JWTRole) IsValid() bool {
	switch r {
	case User, Admin, SuperAdmin:
		return true
	}
	return false
}
