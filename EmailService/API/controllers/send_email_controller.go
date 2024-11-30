package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/send_email"
	commandHandlers "email-service/Application/commmand_handlers/send_email"
	"email-service/Domain/responses"

	"github.com/gin-gonic/gin"
)

type SendEmailController struct {
}

func NewSendEmailController() *SendEmailController {
	return &SendEmailController{}
}

// Handle godoc
// @Summary Send an email
// @Description Sends an email using the provided data.
// @Tags emails
// @Accept json
// @Produce json
// @Param email body contract.SendEmailRequest true "Email data"
// @Success 200 {object} contract.SendEmailResponse200 "Email sent successfully"
// @Failure 400 {object} contract.SendEmailResponse400 "Invalid request format or data"
// @Failure 500 {object} contract.SendEmailResponse500 "Server error during email sending"
// @Router /email-service/api/send-email [post]
func (h *SendEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(contract.SendEmailResponse{
			BaseResponse: responses.NewBaseResponse(400, "Invalid request parameters"),
		})
		return
	}
	command := mappers.MapToSendEmailCommand(&req)
	resp := services.SendToMediator[*commandHandlers.SendEmailCommand, *contract.SendEmailResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *SendEmailController) Route() string {
	return "/send-email"
}

func (h *SendEmailController) Methods() []string {
	return []string{"POST"}
}
