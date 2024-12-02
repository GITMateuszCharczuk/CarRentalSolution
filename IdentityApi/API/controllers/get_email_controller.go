package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/get_email"
	queryHandlers "identity-api/Application/query_handlers/get_email"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetEmailController struct {
	validator *validator.Validate
}

func NewGetEmailController(validator *validator.Validate) *GetEmailController {
	return &GetEmailController{validator: validator}
}

// Handle godoc
// @Summary Get email by ID
// @Description Retrieves an email based on its unique ID.
// @Tags emails
// @Accept json
// @Produce json
// @Param id path string true "Unique Email ID"
// @Success 200 {object} contract.GetEmailResponse200 "Email details retrieved successfully"
// @Failure 400 {object} contract.GetEmailResponse400 "Invalid request parameters"
// @Failure 404 {object} contract.GetEmailResponse404 "Email not found"
// @Failure 500 {object} contract.GetEmailResponse500 "Server error during email retrieval"
// @Router /identity-api/api/emails/{id} [get]
func (h *GetEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	emailID := c.Param("id")
	req := contract.GetEmailRequest{ID: emailID}
	if validateResponse := services.ValidateRequest[contract.GetEmailResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetEmailQuery(&req)
	resp := services.SendToMediator[*queryHandlers.GetEmailQuery, *contract.GetEmailResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetEmailController) Route() string {
	return "/emails/:id"
}

func (h *GetEmailController) Methods() []string {
	return []string{"GET"}
}
