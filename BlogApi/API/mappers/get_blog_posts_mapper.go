package mappers

import (
	contract "identity-api/Application.contract/BlogPosts/GetBlogPosts"
	queries "identity-api/Application/query_handlers/get_blog_posts"
)

func MapToGetBlogPostsQuery(request *contract.GetBlogPostsRequest) queries.GetBlogPostsQuery {
	return queries.GetBlogPostsQuery{
		JwtToken:     request.JwtToken,
		Pagination:   request.Pagination,
		SortQuery:    request.SortQuery,
		Ids:          request.Ids,
		DateTimeFrom: request.DateTimeFrom,
		DateTimeTo:   request.DateTimeTo,
		AuthorIds:    request.AuthorIds,
		Tags:         request.Tags,
		Visible:      request.Visible,
	}
}
