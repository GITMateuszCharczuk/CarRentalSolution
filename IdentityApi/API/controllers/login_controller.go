package controllers

import (
	"identity-api/API/services"
	contract "identity-api/Application.contract/login"
	commands "identity-api/Application/command_handlers/login"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginController struct {
	validator *validator.Validate
}

func NewLoginController(validator *validator.Validate) *LoginController {
	return &LoginController{validator: validator}
}

// Handle godoc
// @Summary User login
// @Description Authenticates a user and returns a JWT token.
// @Tags users
// @Accept json
// @Produce json
// @Param login body contract.LoginRequest true "User login details"
// @Success 200 {object} contract.LoginResponse "Login successful"
// @Failure 400 {object} contract.LoginResponse "Invalid request parameters"
// @Failure 401 {object} contract.LoginResponse "Unauthorized"
// @Failure 500 {object} contract.LoginResponse "Server error during login"
// @Router /identity-api/api/login [post]
func (h *LoginController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(createErrorResponse(400, "Invalid request parameters"))
		return
	}
	if validateResponse := services.ValidateRequest[contract.LoginResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := commands.LoginCommand{
		Email:    req.Email,
		Password: req.Password,
	}
	resp := services.SendToMediator[*commands.LoginCommand, *contract.LoginResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *LoginController) Route() string {
	return "/login"
}

func (h *LoginController) Methods() []string {
	return []string{"POST"}
}
