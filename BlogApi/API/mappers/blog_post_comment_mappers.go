package mappers

import (
	"blog-api/API/services"
	contract "blog-api/Application.contract/BlogPostComments/CreateBlogPostComment"
	delete_contract "blog-api/Application.contract/BlogPostComments/DeketeBlogPostComment"
	get_contract "blog-api/Application.contract/BlogPostComments/GetBlogPostComments"
	commands "blog-api/Application/command_handlers/blog_post_comment/create_blog_post_comment"
	delete_commands "blog-api/Application/command_handlers/blog_post_comment/delete_blog_post_comment"
	get_queries "blog-api/Application/query_handlers/blog_post_comment/get_blog_post_comments"
	"log"
)

func MapToCreateBlogPostCommentCommand(request *contract.CreateBlogPostCommentRequest) commands.CreateBlogPostCommentCommand {
	return commands.CreateBlogPostCommentCommand{
		Description: request.Description,
		BlogPostId:  request.BlogPostId,
		JwtToken:    request.JwtToken,
	}
}

func MapToDeleteBlogPostCommentCommand(request *delete_contract.DeleteBlogPostCommentRequest) delete_commands.DeleteBlogPostCommentCommand {
	return delete_commands.DeleteBlogPostCommentCommand{
		BlogPostCommentId: request.BlogPostCommentId,
		JwtToken:          request.JwtToken,
	}
}

func MapToGetBlogPostCommentsQuery(request *get_contract.GetBlogPostCommentsRequest) get_queries.GetBlogPostCommentsQuery {
	log.Println(len(request.BlogPostIds))
	return get_queries.GetBlogPostCommentsQuery{
		Pagination:   request.Pagination,
		BlogPostIds:  request.BlogPostIds,
		DateTimeFrom: request.DateTimeFrom,
		DateTimeTo:   request.DateTimeTo,
		UserIds:      request.UserIds,
		Sortable:     services.ExtractSorting(request.SortQuery),
	}
}
