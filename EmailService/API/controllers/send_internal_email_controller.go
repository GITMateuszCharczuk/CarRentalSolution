package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/send_internal_email"
	commandHandlers "email-service/Application/commmand_handlers/send_internal_email"
	"email-service/Domain/responses"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SendInternalEmailController struct {
	validator *validator.Validate
}

func NewSendInternalEmailController(validator *validator.Validate) *SendInternalEmailController {
	return &SendInternalEmailController{
		validator: validator,
	}
}

// Handle godoc
// @Summary Send an email
// @Description Sends an email using the provided data.
// @Tags emails
// @Accept json
// @Produce json
// @Param email body contract.SendInternalEmailRequest true "Email data"
// @Success 200 {object} contract.SendInternalEmailResponse200 "Email sent successfully"
// @Failure 400 {object} contract.SendInternalEmailResponse400 "Invalid request format or data"
// @Failure 500 {object} contract.SendInternalEmailResponse500 "Server error during email sending"
// @Router /email-service/api/send-internal-email [post]
func (h *SendInternalEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.SendInternalEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(contract.SendInternalEmailResponse{
			BaseResponse: responses.NewBaseResponse(400, "Invalid request parameters"),
		})
		return
	}
	if validateResponse := services.ValidateRequest[contract.SendInternalEmailResponse](&req, h.validator); validateResponse != nil {
		fmt.Println(validateResponse)
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToSendInternalEmailCommand(&req)
	resp := services.SendToMediator[*commandHandlers.SendInternalEmailCommand, *contract.SendInternalEmailResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *SendInternalEmailController) Route() string {
	return "/send-internal-email"
}

func (h *SendInternalEmailController) Methods() []string {
	return []string{"POST"}
}
