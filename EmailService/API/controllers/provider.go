package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

type Controllers struct {
	All []Controller
}

func NewControllers(all []Controller) *Controllers {
	return &Controllers{All: all}
}

func ProvideControllers(
	getEmailController *GetEmailController,
	getEmailsController *GetEmailsController,
	SendEmailController *SendEmailController,
	SendInternalEmailController *SendInternalEmailController,
) []Controller {
	return []Controller{getEmailController, getEmailsController, SendEmailController, SendInternalEmailController}
}

func ProvideValidator() *validator.Validate {
	return validator.New()
}

var WireSet = wire.NewSet(
	NewGetEmailController,
	NewGetEmailsController,
	NewSendEmailController,
	NewSendInternalEmailController,
	ProvideControllers,
	NewControllers,
	ProvideValidator,
)
