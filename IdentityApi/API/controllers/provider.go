package controllers

import (
	base "identity-api/API/controllers/base"

	"github.com/google/wire"
)

type Controllers struct {
	All []base.Controller
}

func NewControllers(all []base.Controller) *Controllers {
	return &Controllers{All: all}
}

func ProvideControllers(
	getAllUsersController *GetAllUsersController,
	getUserInternalController *GetUserInternalController,
	getUserInfoController *GetUserInfoController,
	registerController *RegisterController,
	loginController *LoginController,
	modifyUserController *ModifyUserController,
	deleteUserController *DeleteUserController,
	validateTokenController *ValidateTokenController,
	refreshTokenController *RefreshTokenController,
) []base.Controller {
	return []base.Controller{
		getAllUsersController,
		getUserInternalController,
		getUserInfoController,
		registerController,
		loginController,
		modifyUserController,
		deleteUserController,
		validateTokenController,
		refreshTokenController,
	}
}

var WireSet = wire.NewSet(
	NewGetAllUsersController,
	NewGetUserInternalController,
	NewGetUserInfoController,
	NewRegisterController,
	NewLoginController,
	NewModifyUserController,
	NewDeleteUserController,
	NewValidateTokenController,
	NewRefreshTokenController,
	ProvideControllers,
	NewControllers,
)
