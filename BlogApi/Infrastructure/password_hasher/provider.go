package password_hasher

import (
	"identity-api/Domain/service_interfaces"

	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
)

func ProvidePasswordHasher() service_interfaces.PasswordHasher {
	return NewBcryptPasswordHasher(bcrypt.DefaultCost)
}

var WireSet = wire.NewSet(ProvidePasswordHasher)
