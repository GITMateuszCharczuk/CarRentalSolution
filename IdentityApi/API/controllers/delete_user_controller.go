package controllers

import (
	"identity-api/API/services"
	contract "identity-api/Application.contract/delete_user"
	commands "identity-api/Application/command_handlers/delete_user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"identity-api/Domain/responses"
)

type DeleteUserController struct {
	validator *validator.Validate
}

func NewDeleteUserController(validator *validator.Validate) *DeleteUserController {
	return &DeleteUserController{validator: validator}
}

// Handle godoc
// @Summary Delete user
// @Description Deletes a user from the system
// @Tags users
// @Accept json
// @Produce json
// @Param delete body contract.DeleteUserRequest true "User deletion details"
// @Success 200 {object} contract.DeleteUserResponse "User deleted successfully"
// @Failure 400 {object} contract.DeleteUserResponse "Invalid request parameters"
// @Failure 401 {object} contract.DeleteUserResponse "Unauthorized"
// @Router /identity-api/api/user [delete]
func (h *DeleteUserController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(responses.NewBaseResponse(400, "Invalid request parameters"))
		return
	}
	if validateResponse := services.ValidateRequest[contract.DeleteUserResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := commands.DeleteUserCommand{
		JwtToken: req.JwtToken,
		ID:       req.ID,
	}
	resp := services.SendToMediator[*commands.DeleteUserCommand, *contract.DeleteUserResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteUserController) Route() string {
	return "/user"
}

func (h *DeleteUserController) Methods() []string {
	return []string{"DELETE"}
}
