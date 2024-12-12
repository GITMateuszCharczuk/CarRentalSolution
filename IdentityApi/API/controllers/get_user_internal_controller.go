package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/get_user_internal"
	queries "identity-api/Application/query_handlers/get_user_internal"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetUserInternalController struct {
	validator *validator.Validate
}

func NewGetUserInternalController(validator *validator.Validate) *GetUserInternalController {
	return &GetUserInternalController{validator: validator}
}

// Handle godoc
// @Summary Get user internal info
// @Description Retrieves user internal info based on the provided token.
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Success 200 {object} contract.GetUserInternalResponse200 "User internal info retrieved successfully"
// @Failure 400 {object} contract.GetUserInternalResponse400 "Invalid request parameters"
// @Failure 404 {object} contract.GetUserInternalResponse404 "User not found"
// @Failure 401 {object} contract.GetUserInternalResponse401 "Unauthorized"
// @Failure 500 {object} contract.GetUserInternalResponse500 "Server error during retrieval"
// @Router /identity-api/api/user/internal [get]
func (h *GetUserInternalController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	req := contract.GetUserInternalRequest{JwtToken: token}
	if validateResponse := services.ValidateRequest[contract.GetUserInternalResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetUserInternalQuery(&req)
	resp := services.SendToMediator[*queries.GetUserInternalQuery, *contract.GetUserInternalResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetUserInternalController) Route() string {
	return "/user/internal"
}

func (h *GetUserInternalController) Methods() []string {
	return []string{"GET"}
}
