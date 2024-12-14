package mappers

import (
	create_contract "identity-api/Application.contract/BlogPostComments/CreateBlogPostComment"
	create_commands "identity-api/Application/command_handlers/create_blog_post_comment"
	delete_commands "identity-api/Application/command_handlers/delete_blog_post_comment"
	get_queries "identity-api/Application/query_handlers/get_blog_post_comments"
)

func MapToCreateBlogPostCommentCommand(request *create_contract.CreateBlogPostCommentRequest) create_commands.CreateBlogPostCommentCommand {
	return create_commands.CreateBlogPostCommentCommand{
		Description: request.Description,
		BlogPostId:  request.BlogPostId,
		JwtToken:    request.JwtToken,
	}
}

func MapToDeleteBlogPostCommentCommand(commentId string, jwtToken string) delete_commands.DeleteBlogPostCommentCommand {
	return delete_commands.DeleteBlogPostCommentCommand{
		BlogPostCommentId: commentId,
		JwtToken:          jwtToken,
	}
}

func MapToGetBlogPostCommentsQuery(blogPostId string, pageSize int, currentPage int) get_queries.GetBlogPostCommentsQuery {
	return get_queries.GetBlogPostCommentsQuery{
		BlogPostId: blogPostId,
		Pagination: struct {
			PageSize    int
			CurrentPage int
		}{
			PageSize:    pageSize,
			CurrentPage: currentPage,
		},
	}
}
