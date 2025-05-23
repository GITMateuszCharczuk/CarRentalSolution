package validators

import (
	models "rental-api/Domain/models/domestic"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func ProvideValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("validRoles", ValidateRoles)
	v.RegisterValidation("validRole", ValidateRole)
	v.RegisterValidation("datetime", validateDateTime)
	v.RegisterValidation("futuredate", validateFutureDate)
	v.RegisterValidation("gtdate", validateGreaterThanDate)
	v.RegisterValidation("validCarOrderStatus", validateCarOrderStatus)
	v.RegisterValidation("validCarOrderStatusArray", validateCarOrderStatusArray)
	RegisterSortQueryValidator[models.CarOfferModel](v, "validCarOfferSortable")
	RegisterSortQueryValidator[models.CarOrderModel](v, "validCarOrderSortable")
	RegisterSortQueryValidator[models.CarOfferTagModel](v, "validCarOfferTagSortable")
	return v
}

var WireSet = wire.NewSet(
	ProvideValidator,
)
