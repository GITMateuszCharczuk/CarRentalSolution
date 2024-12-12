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
// @Description Deletes a user from the system
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "User ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.DeleteUserResponse200 "User deleted successfully"
// @Failure 400 {object} contract.DeleteUserResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.DeleteUserResponse401 "Unauthorized"
// @Failure 404 {object} contract.DeleteUserResponse404 "User not found"
// @Failure 500 {object} contract.DeleteUserResponse500 "Internal server error during deletion"
// @Router /identity-api/api/user/{id} [delete]
func (h *DeleteUserController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	userID := c.Param("id")

	req := contract.DeleteUserRequest{
		JwtToken: token,
		ID:       userID,
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
	return "/user/:id"
}

func (h *DeleteUserController) Methods() []string {
	return []string{"DELETE"}
}
