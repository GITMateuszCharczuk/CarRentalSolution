package controllers

import (
	"identity-api/API/services"
	contract "identity-api/Application.contract/delete_user"
	commands "identity-api/Application/command_handlers/delete_user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteUserController struct {
	validator *validator.Validate
}

func NewDeleteUserController(validator *validator.Validate) *DeleteUserController {
	return &DeleteUserController{validator: validator}
}

// Handle godoc
// @Summary Delete user
// @Description Deletes a user based on the provided token.
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "User token"
// @Success 200 {object} contract.DeleteUserResponse "User deleted successfully"
// @Failure 400 {object} contract.DeleteUserResponse "Invalid request parameters"
// @Failure 404 {object} contract.DeleteUserResponse "User not found"
// @Failure 500 {object} contract.DeleteUserResponse "Server error during deletion"
// @Router /identity-api/api/user/delete [delete]
func (h *DeleteUserController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := c.Query("token")
	req := contract.DeleteUserRequest{Token: token}
	if validateResponse := services.ValidateRequest[contract.DeleteUserResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := commands.DeleteUserCommand{Token: token}
	resp := services.SendToMediator[*commands.DeleteUserCommand, *contract.DeleteUserResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteUserController) Route() string {
	return "/user/delete"
}

func (h *DeleteUserController) Methods() []string {
	return []string{"DELETE"}
}
