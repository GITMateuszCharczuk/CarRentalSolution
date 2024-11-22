package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/send_email"
	command "email-service/Application/commands/send_email"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SendEmailController struct {
	commandHandler *command.SendEmailCommandHandler
}

func NewSendEmailController(handler *command.SendEmailCommandHandler) *SendEmailController {
	return &SendEmailController{
		commandHandler: handler,
	}
}

// Handle godoc
// @Summary Send an email
// @Description Sends an email using the provided data.
// @Tags emails
// @Accept json
// @Produce json
// @Param email body contract.SendEmailRequest true "Email data"
// @Success 200 {object} contract.SendEmailResponse "Email sent successfully"
// @Failure 400 {object} contract.SendEmailResponse "Invalid request format or data"
// @Failure 500 {object} contract.SendEmailResponse "Server error during email sending"
// @Router /send-email [post]
func (h *SendEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	var req contract.SendEmailRequest //bindowanie z taga emaila
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(contract.SendEmailResponse{
			Title:   "StatusBadRequest",
			Message: fmt.Sprintf("Bad input data: %v", err),
		})
		return
	}

	command := mappers.MapToSendEmailCommand(&req)
	resp, err := h.commandHandler.Execute(command)
	if err != nil {
		responseSender.Send(contract.SendEmailResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *SendEmailController) Route() string {
	return "/send-email"
}

func (h *SendEmailController) Methods() []string {
	return []string{"POST"}
}
