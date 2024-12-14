package validators

import (
	models "identity-api/Domain/models/domestic"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func ProvideValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("validRoles", ValidateRoles)
	v.RegisterValidation("validRole", ValidateRole)
	RegisterSortQueryValidator[models.BlogPostCommentModel](v, "validCommentSortable")
	RegisterSortQueryValidator[models.BlogPostTagModel](v, "validTagSortable")
	RegisterSortQueryValidator[models.BlogPostRequestModel](v, "validBlogPostSortable")
	return v
}

var WireSet = wire.NewSet(
	ProvideValidator,
)
