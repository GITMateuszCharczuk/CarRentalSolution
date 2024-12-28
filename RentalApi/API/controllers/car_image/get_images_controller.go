package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/images/get_all_images"
	queries "rental-api/Application/query_handlers/car_image/get_images"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetImagesController struct {
	validator *validator.Validate
}

func NewGetImagesController(validator *validator.Validate) *GetImagesController {
	return &GetImagesController{validator: validator}
}

// Handle godoc
// @Summary Get all images for a car offer
// @Description Retrieves all images for a specific car offer
// @Tags car-images
// @Accept json
// @Produce json
// @Param carOfferId path string true "Car Offer ID" example:"371eb93a-054c-44db-b429-0a0ebe87c3b9"
// @Success 200 {object} contract.GetAllImagesResponse200 "Images retrieved successfully"
// @Failure 400 {object} contract.GetAllImagesResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.GetAllImagesResponse401 "Unauthorized"
// @Failure 403 {object} contract.GetAllImagesResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.GetAllImagesResponse404 "Image or car offer not found"
// @Failure 500 {object} contract.GetAllImagesResponse500 "Server error during retrieval"
// @Router /rental-api/api/car-offers/images/{carOfferId} [get]
func (h *GetImagesController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetAllImagesRequest{
		CarOfferId: services.ExtractFromPath(c, "carOfferId"),
	}
	if validateResponse := services.ValidateRequest[contract.GetAllImagesResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToGetImagesQuery(&req)
	resp := services.SendToMediator[*queries.GetImagesQuery, *contract.GetAllImagesResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *GetImagesController) Route() string {
	return "/car-offers/images/:carOfferId"
}

func (h *GetImagesController) Methods() []string {
	return []string{"GET"}
}
