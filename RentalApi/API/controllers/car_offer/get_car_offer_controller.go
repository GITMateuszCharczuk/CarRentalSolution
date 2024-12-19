package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_offers/get_car_offer"
	queries "rental-api/Application/query_handlers/car_offer/get_car_offer"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetCarOfferController struct {
	validator *validator.Validate
}

func NewGetCarOfferController(validator *validator.Validate) *GetCarOfferController {
	return &GetCarOfferController{validator: validator}
}

// Handle godoc
// @Summary Get a car offer by ID
// @Description Retrieves a specific car offer by its ID
// @Tags car-offers
// @Accept json
// @Produce json
// @Param id path string true "Car Offer ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.GetCarOfferResponse200 "Car offer retrieved successfully"
// @Failure 404 {object} contract.GetCarOfferResponse404 "Car offer not found"
// @Failure 500 {object} contract.GetCarOfferResponse500 "Server error during retrieval"
// @Router /rental-api/api/car-offers/{id} [get]
func (h *GetCarOfferController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetCarOfferRequest{
		CarOfferId: services.ExtractFromPath(c, "id"),
	}
	if validateResponse := services.ValidateRequest[contract.GetCarOfferResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetCarOfferQuery(&req)
	resp := services.SendToMediator[*queries.GetCarOfferQuery, *contract.GetCarOfferResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetCarOfferController) Route() string {
	return "/car-offers/:id"
}

func (h *GetCarOfferController) Methods() []string {
	return []string{"GET"}
}
