package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/get_email"
	queryHandlers "email-service/Application/query_handlers/get_email"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

type GetEmailController struct {
}

func NewGetEmailController() *GetEmailController {
	return &GetEmailController{}
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
// @Router /email-service/api/emails/{id} [get]
func (h *GetEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	emailID := c.Param("id")

	req := contract.GetEmailRequest{ID: emailID}
	query := mappers.MapToGetEmailQuery(&req)
	resp, err := mediatr.Send[*queryHandlers.GetEmailQuery, *contract.GetEmailResponse](c.Request.Context(), &query)
	if err != nil {
		responseSender.Send(contract.GetEmailResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *GetEmailController) Route() string {
	return "/emails/:id"
}

func (h *GetEmailController) Methods() []string {
	return []string{"GET"}
}
