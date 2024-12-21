package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/get_emails"
	queries "email-service/Application/query_handlers/get_emails"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetEmailsController struct {
	validator *validator.Validate
}

func NewGetEmailsController(validator *validator.Validate) *GetEmailsController {
	return &GetEmailsController{validator: validator}
}

// Handle godoc
// @Summary Get all emails
// @Description Retrieves a list of all emails.
// @Tags emails
// @Accept json
// @Produce json
// @Param token query string true "JWT Token"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of emails per page" default(10)
// @Success 200 {object} contract.GetEmailsResponse200 "List of emails retrieved successfully"
// @Failure 400 {object} contract.GetEmailsResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.GetEmailsResponse401 "Unauthorized"
// @Failure 403 {object} contract.GetEmailsResponse403 "Forbidden"
// @Failure 500 {object} contract.GetEmailsResponse500 "Server error during emails retrieval"
// @Router /email-service/api/emails [get]
func (h *GetEmailsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetEmailsRequest{
		JwtToken:   services.GetJwtTokenFromQuery(c),
		Pagination: services.ExtractPagination(c),
	}
	if validateResponse := services.ValidateRequest[contract.GetEmailsResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetEmailsQuery(&req)
	resp := services.SendToMediator[*queries.GetEmailsQuery, *contract.GetEmailsResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetEmailsController) Route() string {
	return "/emails"
}

func (h *GetEmailsController) Methods() []string {
	return []string{"GET"}
}
