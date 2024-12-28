package queries

import (
	"context"
	contract "rental-api/Application.contract/images/get_all_images"
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_image_repository"
	"rental-api/Domain/responses"
)

type GetImagesQueryHandler struct {
	carImageQueryRepository repository_interfaces.CarImageQueryRepository
}

func NewGetImagesQueryHandler(
	carImageQueryRepository repository_interfaces.CarImageQueryRepository,
) *GetImagesQueryHandler {
	return &GetImagesQueryHandler{
		carImageQueryRepository: carImageQueryRepository,
	}
}

func (h *GetImagesQueryHandler) Handle(ctx context.Context, query *GetImagesQuery) (*contract.GetAllImagesResponse, error) {
	images, err := h.carImageQueryRepository.GetImagesByCarOfferId(query.CarOfferId)
	if err != nil {
		response := responses.NewResponse[contract.GetAllImagesResponse](500, "Failed to retrieve images")
		return &response, nil
	}

	if images == nil {
		images = &[]models.CarOfferImageModel{}
	}

	return &contract.GetAllImagesResponse{
		BaseResponse: responses.NewBaseResponse(200, "Images retrieved successfully"),
		Items:        *images,
	}, nil
}
