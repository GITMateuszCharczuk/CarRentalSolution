package controllers

import (
	mappers "identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/register"
	commands "identity-api/Application/command_handlers/register"
	"identity-api/Domain/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterController struct {
	validator *validator.Validate
}

func NewRegisterController(validator *validator.Validate) *RegisterController {
	return &RegisterController{validator: validator}
}

// Handle godoc
// @Summary Register a new user
// @Description Registers a new user with the provided details.
// @Tags users
// @Accept json
// @Produce json
// @Param register body contract.RegisterUserRequest true "User registration details"
// @Success 201 {object} contract.RegisterUserResponse201 "User registered successfully"
// @Failure 400 {object} contract.RegisterUserResponse400 "Invalid request parameters"
// @Failure 500 {object} contract.RegisterUserResponse500 "Server error during registration"
// @Router /identity-api/api/register [post]
func (h *RegisterController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := responses.NewResponse[contract.RegisterUserResponse](400, "Invalid request parameters")
		responseSender.Send(response)
		return
	}
	if validateResponse := services.ValidateRequest[contract.RegisterUserResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToRegisterCommand(&req)
	resp := services.SendToMediator[*commands.RegisterUserCommand, *contract.RegisterUserResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *RegisterController) Route() string {
	return "/register"
}

func (h *RegisterController) Methods() []string {
	return []string{"POST"}
}
