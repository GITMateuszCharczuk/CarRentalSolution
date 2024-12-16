// API/mappers/blog_post_like_mappers.go
package mappers

import (
	contract "identity-api/Application.contract/BlogPostLikes/CreateLikeForBlogPost"
	delete_contract "identity-api/Application.contract/BlogPostLikes/DeleteLikeForBlogPost"
	get_contract "identity-api/Application.contract/BlogPostLikes/GetLikesForBlogPost"
	commands "identity-api/Application/command_handlers/blog_post_like/create_like_for_blog_post"
	delete_commands "identity-api/Application/command_handlers/blog_post_like/delete_like_for_blog_post"
	get_queries "identity-api/Application/query_handlers/blog_post_like/get_likes_for_blog_post"
)

func MapToCreateLikeForBlogPostCommand(request *contract.CreateLikeForBlogPostRequest) commands.CreateLikeForBlogPostCommand {
	return commands.CreateLikeForBlogPostCommand{
		BlogPostId: request.BlogPostId,
		JwtToken:   request.JwtToken,
	}
}

func MapToDeleteLikeForBlogPostCommand(request *delete_contract.DeleteLikeForBlogPostRequest) delete_commands.DeleteLikeForBlogPostCommand {
	return delete_commands.DeleteLikeForBlogPostCommand{
		BlogPostId: request.BlogPostId,
		JwtToken:   request.JwtToken,
	}
}

func MapToGetLikesForBlogPostQuery(request *get_contract.GetLikesForBlogPostRequest) get_queries.GetLikesForBlogPostQuery {
	return get_queries.GetLikesForBlogPostQuery{
		BlogPostId: request.BlogPostId,
	}
}
