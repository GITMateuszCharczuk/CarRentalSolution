package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/get_user_id"
	queries "identity-api/Application/query_handlers/get_user_id"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetUserIDController struct {
	validator *validator.Validate
}

func NewGetUserIDController(validator *validator.Validate) *GetUserIDController {
	return &GetUserIDController{validator: validator}
}

// Handle godoc
// @Summary Get user ID
// @Description Retrieves user ID based on the provided token.
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Success 200 {object} contract.GetUserIDResponse200 "User ID retrieved successfully"
// @Failure 400 {object} contract.GetUserIDResponse400 "Invalid request parameters"
// @Failure 404 {object} contract.GetUserIDResponse404 "User not found"
// @Failure 401 {object} contract.GetUserIDResponse401 "Unauthorized"
// @Failure 500 {object} contract.GetUserIDResponse500 "Server error during retrieval"
// @Router /identity-api/api/user/id [get]
func (h *GetUserIDController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	req := contract.GetUserIDRequest{JwtToken: token}
	if validateResponse := services.ValidateRequest[contract.GetUserIDResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetUserIDQuery(&req)
	resp := services.SendToMediator[*queries.GetUserIDQuery, *contract.GetUserIDResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetUserIDController) Route() string {
	return "/user/id"
}

func (h *GetUserIDController) Methods() []string {
	return []string{"GET"}
}
