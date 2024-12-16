package mappers

import (
	"identity-api/API/services"
	create_contract "identity-api/Application.contract/BlogPosts/CreateBlogPost"
	delete_contract "identity-api/Application.contract/BlogPosts/DeleteBlogPost"
	get_blog_post_contract "identity-api/Application.contract/BlogPosts/GetBlogPost"
	get_blog_posts_contract "identity-api/Application.contract/BlogPosts/GetBlogPosts"
	update_contract "identity-api/Application.contract/BlogPosts/UpdateBlogPost"
	commands "identity-api/Application/command_handlers/blog_post/create_blog_post"
	delete_commands "identity-api/Application/command_handlers/blog_post/delete_blog_post"
	update_commands "identity-api/Application/command_handlers/blog_post/update_blog_post"
	get_blog_post_queries "identity-api/Application/query_handlers/blog_post/get_blog_post"
	get_blog_posts_queries "identity-api/Application/query_handlers/blog_post/get_blog_posts"
)

func MapToCreateBlogPostCommand(request *create_contract.CreateBlogPostRequest) commands.CreateBlogPostCommand {
	return commands.CreateBlogPostCommand{
		Heading:          request.Heading,
		PageTitle:        request.PageTitle,
		Content:          request.Content,
		ShortDescription: request.ShortDescription,
		FeaturedImageUrl: request.FeaturedImageUrl,
		UrlHandle:        request.UrlHandle,
		Visible:          request.Visible,
		Tags:             request.Tags,
		JwtToken:         request.JwtToken,
	}
}

func MapToUpdateBlogPostCommand(request *update_contract.UpdateBlogPostRequest) update_commands.UpdateBlogPostCommand {
	return update_commands.UpdateBlogPostCommand{
		Id:               request.Id,
		Heading:          request.Heading,
		PageTitle:        request.PageTitle,
		Content:          request.Content,
		ShortDescription: request.ShortDescription,
		FeaturedImageUrl: request.FeaturedImageUrl,
		UrlHandle:        request.UrlHandle,
		PublishedDate:    request.PublishedDate,
		Author:           request.Author,
		Visible:          request.Visible,
		Tags:             request.Tags,
		JwtToken:         request.JwtToken,
	}
}

func MapToDeleteBlogPostCommand(request *delete_contract.DeleteBlogPostRequest) delete_commands.DeleteBlogPostCommand {
	return delete_commands.DeleteBlogPostCommand{
		ID:       request.BlogPostId,
		JwtToken: request.JwtToken,
	}
}

func MapToGetBlogPostsQuery(request *get_blog_posts_contract.GetBlogPostsRequest) get_blog_posts_queries.GetBlogPostsQuery {
	return get_blog_posts_queries.GetBlogPostsQuery{
		Ids:          request.Ids,
		DateTimeFrom: request.DateTimeFrom,
		DateTimeTo:   request.DateTimeTo,
		AuthorIds:    request.AuthorIds,
		Tags:         request.Tags,
		Visible:      request.Visible,
		Sortable:     services.ExtractSorting(request.SortQuery),
		Pagination:   request.Pagination,
	}
}

func MapToGetBlogPostQuery(request *get_blog_post_contract.GetBlogPostRequest) get_blog_post_queries.GetBlogPostQuery {
	return get_blog_post_queries.GetBlogPostQuery{
		ID: request.BlogPostId,
	}
}
