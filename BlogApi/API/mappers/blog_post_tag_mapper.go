package mappers

import (
	"blog-api/API/services"
	get_tags_contract "blog-api/Application.contract/Tags/GetTags"
	get_tags_queries "blog-api/Application/query_handlers/blog_post_tag/get_tags"
	"log"
)

func MapToGetTagsQuery(request *get_tags_contract.GetTagsRequest) get_tags_queries.GetTagsQuery {
	log.Println(request.BlogPostId)
	return get_tags_queries.GetTagsQuery{
		BlogPostId: request.BlogPostId,
		Sortable:   services.ExtractSorting(request.SortQuery),
	}
}
