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
// @Description Retrieves a list of all users with optional pagination and sorting.
// @Tags users
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param page_size query int false "Page size" example:"10"
// @Param current_page query int false "Current page" example:"1"
// @Param sort_fields query []string false "Sort fields (format: field:direction)" example:"name:asc,email:desc"
// @Success 200 {object} contract.GetAllUsersResponse200 "Users retrieved successfully"
// @Failure 400 {object} contract.GetAllUsersResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.GetAllUsersResponse401 "Unauthorized"
// @Failure 403 {object} contract.GetAllUsersResponse403 "Insufficient privileges"
// @Failure 500 {object} contract.GetAllUsersResponse500 "Server error during retrieval"
// @Router /identity-api/api/users [get]
func (h *GetAllUsersController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	req := contract.GetAllUsersRequest{
		JwtToken:   services.GetJwtTokenFromQuery(c),
		Pagination: services.ExtractPagination(c),
		SortQuery:  services.ExtractSortQuery(c),
	}

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
