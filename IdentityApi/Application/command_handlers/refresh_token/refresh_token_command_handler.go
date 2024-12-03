package commands

import (
	"context"
	contract "identity-api/Application.contract/refresh_token"
	models "identity-api/Domain/models/token"
	"identity-api/Domain/responses"
	service_interfaces "identity-api/Domain/service_interfaces"
)

type RefreshTokenCommandHandler struct {
	tokenService service_interfaces.JWTTokenService
}

func NewRefreshTokenCommandHandler(tokenService service_interfaces.JWTTokenService) *RefreshTokenCommandHandler {
	return &RefreshTokenCommandHandler{tokenService: tokenService}
}

func (h *RefreshTokenCommandHandler) Handle(ctx context.Context, command *RefreshTokenCommand) (*contract.RefreshTokenResponse, error) {
	constructedRefreshToken := models.NewRefreshToken(command.RefreshToken)
	token, err := h.tokenService.RefreshToken(constructedRefreshToken)
	if err != nil {
		return &contract.RefreshTokenResponse{
			BaseResponse: responses.NewBaseResponse(401, "Invalid refresh token"),
		}, nil
	}

	return &contract.RefreshTokenResponse{
		BaseResponse: responses.NewBaseResponse(200, "Token refreshed successfully"),
		JwtToken:     token,
	}, nil
}
