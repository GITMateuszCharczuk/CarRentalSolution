package mappers

import (
	"identity-api/API/services"
	get_tags_contract "identity-api/Application.contract/Tags/GetTags"
	get_tags_queries "identity-api/Application/query_handlers/blog_post_tag/get_tags"
)

func MapToGetTagsQuery(request *get_tags_contract.GetTagsRequest) get_tags_queries.GetTagsQuery {
	return get_tags_queries.GetTagsQuery{
		BlogPostId: request.BlogPostId,
		Sortable:   services.ExtractSorting(request.SortQuery),
	}
}
