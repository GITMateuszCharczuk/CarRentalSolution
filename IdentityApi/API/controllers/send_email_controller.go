package controllers

import (
	"fmt"
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/send_email"
	commandHandlers "identity-api/Application/commmand_handlers/send_email"
	"identity-api/Domain/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SendEmailController struct {
	validator *validator.Validate
}

func NewSendEmailController(validator *validator.Validate) *SendEmailController {
	return &SendEmailController{
		validator: validator,
	}
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
// @Router /identity-api/api/send-email [post]
func (h *SendEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	var req contract.SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(contract.SendEmailResponse{
			BaseResponse: responses.NewBaseResponse(400, "Invalid request parameters"),
		})
		return
	}
	if validateResponse := services.ValidateRequest[contract.SendEmailResponse](&req, h.validator); validateResponse != nil {
		fmt.Println(validateResponse)
		responseSender.Send(validateResponse)
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
