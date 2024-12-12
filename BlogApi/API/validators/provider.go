package validators

import (
	models "identity-api/Domain/models/user"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func ProvideValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("validRoles", ValidateRoles)
	v.RegisterValidation("validRole", ValidateRole)
	RegisterSortQueryValidator[models.UserSecureInfo](v, "validUserSortable")
	RegisterSortQueryValidator[models.BlogPostComment](v, "validCommentSortable")
	return v
}

var WireSet = wire.NewSet(
	ProvideValidator,
)
