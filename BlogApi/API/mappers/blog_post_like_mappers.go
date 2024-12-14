package mappers

import (
	create_commands "identity-api/Application/command_handlers/create_like_for_blog_post"
	delete_commands "identity-api/Application/command_handlers/delete_like_for_blog_post"
	get_queries "identity-api/Application/query_handlers/get_likes_for_blog_post"
)

func MapToCreateLikeForBlogPostCommand(blogPostId string, jwtToken string) create_commands.CreateLikeForBlogPostCommand {
	return create_commands.CreateLikeForBlogPostCommand{
		BlogPostId: blogPostId,
		JwtToken:   jwtToken,
	}
}

func MapToDeleteLikeForBlogPostCommand(blogPostId string, jwtToken string) delete_commands.DeleteLikeForBlogPostCommand {
	return delete_commands.DeleteLikeForBlogPostCommand{
		BlogPostId: blogPostId,
		JwtToken:   jwtToken,
	}
}

func MapToGetLikesForBlogPostQuery(blogPostId string) get_queries.GetLikesForBlogPostQuery {
	return get_queries.GetLikesForBlogPostQuery{
		BlogPostId: blogPostId,
	}
}
