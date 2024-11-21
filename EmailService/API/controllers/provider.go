package controllers

import "github.com/google/wire"

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
) []Controller {
	return []Controller{getEmailController, getEmailsController, SendEmailController}
}

var WireSet = wire.NewSet(
	NewGetEmailController,
	NewGetEmailsController,
	NewSendEmailController,
	ProvideControllers,
	NewControllers,
)
