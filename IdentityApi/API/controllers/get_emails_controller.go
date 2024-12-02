package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/get_emails"
	queries "identity-api/Application/query_handlers/get_emails"
	"identity-api/Domain/responses"

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
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of emails per page" default(10)
// @Success 200 {object} contract.GetEmailsResponse200 "List of emails retrieved successfully"
// @Failure 400 {object} contract.GetEmailsResponse400 "Invalid request parameters"
// @Failure 500 {object} contract.GetEmailsResponse500 "Server error during emails retrieval"
// @Router /identity-api/api/emails [get]
func (h *GetEmailsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.GetEmailsRequest
	if err := c.ShouldBindQuery(&req.Pagination); err != nil {
		responseSender.Send(contract.GetEmailsResponse{
			BaseResponse: responses.NewBaseResponse(400, "Invalid request parameters"),
		})
		return
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
