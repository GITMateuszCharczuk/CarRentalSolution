package queries

import (
	"context"
	contract "rental-api/Application.contract/tags/get_tags"
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_tag_repository"
	"rental-api/Domain/responses"
)

type GetTagsQueryHandler struct {
	carTagQueryRepository repository_interfaces.CarTagQueryRepository
}

func NewGetTagsQueryHandler(
	carTagQueryRepository repository_interfaces.CarTagQueryRepository,
) *GetTagsQueryHandler {
	return &GetTagsQueryHandler{
		carTagQueryRepository: carTagQueryRepository,
	}
}

func (h *GetTagsQueryHandler) Handle(ctx context.Context, query *GetTagsQuery) (*contract.GetTagsResponse, error) {
	tags, err := h.carTagQueryRepository.GetTags(
		&query.Sortable,
		query.CarOfferId,
	)

	if err != nil {
		response := responses.NewResponse[contract.GetTagsResponse](500, "Failed to retrieve tags")
		return &response, nil
	}

	if tags == nil {
		tags = &[]models.CarOfferTagModel{}
	}

	return &contract.GetTagsResponse{
		BaseResponse: responses.NewBaseResponse(200, "Tags retrieved successfully"),
		Items:        *tags,
	}, nil
}
