package controllers

import (
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
// @Param modify body contract.ModifyUserRequest true "User modification details"
// @Success 200 {object} contract.ModifyUserResponse "User modified successfully"
// @Failure 400 {object} contract.ModifyUserResponse "Invalid request parameters"
// @Failure 401 {object} contract.ModifyUserResponse "Unauthorized"
// @Router /identity-api/api/user [put]
func (h *ModifyUserController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.ModifyUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(responses.NewBaseResponse(400, "Invalid request parameters"))
		return
	}
	if validateResponse := services.ValidateRequest[contract.ModifyUserResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToModifyUserCommand(&req)
	resp := services.SendToMediator[*commands.ModifyUserCommand, *contract.ModifyUserResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *ModifyUserController) Route() string {
	return "/user"
}

func (h *ModifyUserController) Methods() []string {
	return []string{"PUT"}
}
