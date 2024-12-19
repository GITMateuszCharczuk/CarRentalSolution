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
	tags, err := h.carTagQueryRepository.GetTagsByCarOfferId(
		query.CarOfferId,
		&query.Sortable,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			response := responses.NewResponse[contract.GetTagsResponse](200, "No tags found")
			return &response, nil
		}
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
