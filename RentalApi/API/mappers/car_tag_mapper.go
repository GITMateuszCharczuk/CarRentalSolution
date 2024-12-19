package mappers

import (
	"rental-api/API/services"
	get_tags_contract "rental-api/Application.contract/tags/get_tags"
	get_tags_queries "rental-api/Application/query_handlers/car_tag/get_tags"
)

func MapToGetTagsQuery(request *get_tags_contract.GetTagsRequest) get_tags_queries.GetTagsQuery {
	return get_tags_queries.GetTagsQuery{
		CarOfferId: request.CarOfferId,
		Sortable:   services.ExtractSorting(request.SortQuery),
	}
}
