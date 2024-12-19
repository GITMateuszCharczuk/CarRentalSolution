package queries

import (
	"context"
	contract "rental-api/Application.contract/car_offers/get_car_offers"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type GetCarOffersQueryHandler struct {
	carOfferQueryRepository repository_interfaces.CarOfferQueryRepository
	dataFetcher             data_fetcher.MicroserviceConnector
}

func NewGetCarOffersQueryHandler(
	carOfferQueryRepository repository_interfaces.CarOfferQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *GetCarOffersQueryHandler {
	return &GetCarOffersQueryHandler{
		carOfferQueryRepository: carOfferQueryRepository,
		dataFetcher:             dataFetcher,
	}
}

func (h *GetCarOffersQueryHandler) Handle(ctx context.Context, query *GetCarOffersQuery) (*contract.GetCarOffersResponse, error) {
	result, err := h.carOfferQueryRepository.GetCarOffers(
		&query.Pagination,
		&query.Sortable,
		query.Ids,
		query.DateTimeFrom,
		query.DateTimeTo,
		query.Tags,
		query.Visible,
	)

	if err != nil {
		response := responses.NewResponse[contract.GetCarOffersResponse](500, "Failed to retrieve car offers")
		return &response, nil
	}

	return &contract.GetCarOffersResponse{
		BaseResponse:    responses.NewBaseResponse(200, "Car offers retrieved successfully"),
		PaginatedResult: *result,
	}, nil
}
