package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/get_user_info"
	queries "identity-api/Application/query_handlers/get_user_info"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetUserInfoController struct {
	validator *validator.Validate
}

func NewGetUserInfoController(validator *validator.Validate) *GetUserInfoController {
	return &GetUserInfoController{validator: validator}
}

// Handle godoc
// @Summary Get user info
// @Description Retrieves user information based on the provided token and user ID.
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id query string false "User ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.GetUserInfoResponse "User info retrieved successfully"
// @Failure 400 {object} contract.GetUserInfoResponse "Invalid request parameters"
// @Failure 404 {object} contract.GetUserInfoResponse "User not found"
// @Failure 500 {object} contract.GetUserInfoResponse "Server error during retrieval"
// @Router /identity-api/api/user/info [get]
func (h *GetUserInfoController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	userID := c.Query("id")

	req := contract.GetUserInfoRequest{
		JwtToken: token,
		Id:       userID,
	}

	if validateResponse := services.ValidateRequest[contract.GetUserInfoResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	query := mappers.MapToGetUserInfoQuery(&req)
	resp := services.SendToMediator[*queries.GetUserInfoQuery, *contract.GetUserInfoResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetUserInfoController) Route() string {
	return "/user/info"
}

func (h *GetUserInfoController) Methods() []string {
	return []string{"GET"}
}
