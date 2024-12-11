package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/validate_token"
	commands "identity-api/Application/command_handlers/validate_token"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateTokenController struct {
	validator *validator.Validate
}

func NewValidateTokenController(validator *validator.Validate) *ValidateTokenController {
	return &ValidateTokenController{validator: validator}
}

// Handle godoc
// @Summary Validate token
// @Description Validates the provided JWT token and returns user roles.
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Success 200 {object} contract.ValidateTokenResponse200 "Token is valid"
// @Failure 400 {object} contract.ValidateTokenResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.ValidateTokenResponse401 "Unauthorized"
// @Failure 500 {object} contract.ValidateTokenResponse500 "Server error during validation"
// @Router /identity-api/api/token/validate [get]
func (h *ValidateTokenController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	req := contract.ValidateTokenRequest{JwtToken: token}
	if validateResponse := services.ValidateRequest[contract.ValidateTokenResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToValidateTokenCommand(&req)
	resp := services.SendToMediator[*commands.ValidateTokenCommand, *contract.ValidateTokenResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *ValidateTokenController) Route() string {
	return "/token/validate"
}

func (h *ValidateTokenController) Methods() []string {
	return []string{"GET"}
}
