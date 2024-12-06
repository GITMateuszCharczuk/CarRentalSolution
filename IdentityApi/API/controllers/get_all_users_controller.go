package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/get_all_users"
	queries "identity-api/Application/query_handlers/get_all_users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetAllUsersController struct {
	validator *validator.Validate
}

func NewGetAllUsersController(validator *validator.Validate) *GetAllUsersController {
	return &GetAllUsersController{validator: validator}
}

// Handle godoc
// @Summary Get all users
// @Description Retrieves a list of all users.
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Success 200 {object} contract.GetAllUsersResponse "Users retrieved successfully"
// @Failure 400 {object} contract.GetAllUsersResponse "Invalid request parameters"
// @Failure 500 {object} contract.GetAllUsersResponse "Server error during retrieval"
// @Router /identity-api/api/users [get]
func (h *GetAllUsersController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	req := contract.GetAllUsersRequest{JwtToken: token}
	if validateResponse := services.ValidateRequest[contract.GetAllUsersResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetAllUsersQuery(&req)
	resp := services.SendToMediator[*queries.GetAllUsersQuery, *contract.GetAllUsersResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetAllUsersController) Route() string {
	return "/users"
}

func (h *GetAllUsersController) Methods() []string {
	return []string{"GET"}
}
