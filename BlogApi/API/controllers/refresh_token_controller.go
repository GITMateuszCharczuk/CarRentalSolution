package controllers

import (
	"identity-api/API/services"
	contract "identity-api/Application.contract/refresh_token"
	commands "identity-api/Application/command_handlers/refresh_token"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RefreshTokenController struct {
	validator *validator.Validate
}

func NewRefreshTokenController(validator *validator.Validate) *RefreshTokenController {
	return &RefreshTokenController{validator: validator}
}

// Handle godoc
// @Summary Refresh token
// @Description Refreshes the JWT token using a refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param token query string true "Refresh token" example:"your.refresh.token.here"
// @Success 200 {object} contract.RefreshTokenResponse200 "Token refreshed successfully"
// @Failure 400 {object} contract.RefreshTokenResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.RefreshTokenResponse401 "Invalid refresh token"
// @Failure 500 {object} contract.RefreshTokenResponse500 "Server error during token refresh"
// @Router /identity-api/api/token/refresh [post]
func (h *RefreshTokenController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtRefreshTokenFromQuery(c)
	req := contract.RefreshTokenRequest{JwtRefreshToken: token}
	if validateResponse := services.ValidateRequest[contract.RefreshTokenResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := commands.RefreshTokenCommand{JwtRefreshToken: req.JwtRefreshToken}
	resp := services.SendToMediator[*commands.RefreshTokenCommand, *contract.RefreshTokenResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *RefreshTokenController) Route() string {
	return "/token/refresh"
}

func (h *RefreshTokenController) Methods() []string {
	return []string{"POST"}
}
