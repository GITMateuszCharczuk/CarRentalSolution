package controllers

import (
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
// @Param token query string true "User token"
// @Success 200 {object} contract.GetUserIDResponse "User ID retrieved successfully"
// @Failure 400 {object} contract.GetUserIDResponse "Invalid request parameters"
// @Failure 404 {object} contract.GetUserIDResponse "User not found"
// @Failure 500 {object} contract.GetUserIDResponse "Server error during retrieval"
// @Router /identity-api/api/user/id [get]
func (h *GetUserIDController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := c.Query("token")
	req := contract.GetUserIDRequest{Token: token}
	if validateResponse := services.ValidateRequest[contract.GetUserIDResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := queries.GetUserIDQuery{Token: token}
	resp := services.SendToMediator[*queries.GetUserIDQuery, *contract.GetUserIDResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetUserIDController) Route() string {
	return "/user/id"
}

func (h *GetUserIDController) Methods() []string {
	return []string{"GET"}
}
