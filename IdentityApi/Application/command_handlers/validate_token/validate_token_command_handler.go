package commands

import (
	"context"
	contract "identity-api/Application.contract/validate_token"
	"identity-api/Application/services"
	"identity-api/Domain/responses"
	"identity-api/Domain/service_interfaces"
)

type ValidateTokenCommandHandler struct {
	tokenService service_interfaces.JWTTokenService
}

func NewValidateTokenCommandHandler(tokenService service_interfaces.JWTTokenService) *ValidateTokenCommandHandler {
	return &ValidateTokenCommandHandler{tokenService: tokenService}
}

func (h *ValidateTokenCommandHandler) Handle(ctx context.Context, command *ValidateTokenCommand) (*contract.ValidateTokenResponse, error) {
	_, roles, err := h.tokenService.ValidateToken(command.JwtToken)
	if err != nil {
		return &contract.ValidateTokenResponse{
			BaseResponse: responses.NewBaseResponse(401, "Invalid token"),
			Valid:        false,
		}, nil
	}

	return &contract.ValidateTokenResponse{
		BaseResponse: responses.NewBaseResponse(200, "Token is valid"),
		Valid:        true,
		Roles:        services.ConvertRolesToString(roles),
	}, nil
}
