package controllers

import (
	"fmt"
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/modify_user"
	commands "identity-api/Application/command_handlers/modify_user"

	"identity-api/Domain/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ModifyUserController struct {
	validator *validator.Validate
}

func NewModifyUserController(validator *validator.Validate) *ModifyUserController {
	return &ModifyUserController{validator: validator}
}

// Handle godoc
// @Summary Modify user
// @Description Modifies user information
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param modify body contract.ModifyUserRequest true "User modification details"
// @Success 200 {object} contract.ModifyUserResponse200 "User modified successfully"
// @Failure 400 {object} contract.ModifyUserResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.ModifyUserResponse401 "Unauthorized"
// @Failure 404 {object} contract.ModifyUserResponse404 "User not found"
// @Failure 500 {object} contract.ModifyUserResponse500 "Server error during modification"
// @Router /identity-api/api/user [put]
func (h *ModifyUserController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	var req contract.ModifyUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		response := responses.NewResponse[contract.ModifyUserResponse](400, "Invalid request parameters")
		responseSender.Send(response)
		return
	}
	if validateResponse := services.ValidateRequest[contract.ModifyUserResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToModifyUserCommand(&req)
	command.JwtToken = token
	resp := services.SendToMediator[*commands.ModifyUserCommand, *contract.ModifyUserResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *ModifyUserController) Route() string {
	return "/user"
}

func (h *ModifyUserController) Methods() []string {
	return []string{"PUT"}
}
