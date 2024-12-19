package queries

import (
	"context"
	contract "rental-api/Application.contract/car_offers/get_car_offer"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/responses"
)

type GetCarOfferQueryHandler struct {
	carOfferQueryRepository repository_interfaces.CarOfferQueryRepository
}

func NewGetCarOfferQueryHandler(
	carOfferQueryRepository repository_interfaces.CarOfferQueryRepository,
) *GetCarOfferQueryHandler {
	return &GetCarOfferQueryHandler{
		carOfferQueryRepository: carOfferQueryRepository,
	}
}

func (h *GetCarOfferQueryHandler) Handle(ctx context.Context, query *GetCarOfferQuery) (*contract.GetCarOfferResponse, error) {
	carOffer, err := h.carOfferQueryRepository.GetCarOfferByID(query.ID)
	if err != nil {
		response := responses.NewResponse[contract.GetCarOfferResponse](500, "Failed to retrieve car offer")
		return &response, nil
	}

	if carOffer == nil {
		response := responses.NewResponse[contract.GetCarOfferResponse](404, "Car offer not found")
		return &response, nil
	}

	return &contract.GetCarOfferResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car offer retrieved successfully"),
		CarOffer:     *carOffer,
	}, nil
}
