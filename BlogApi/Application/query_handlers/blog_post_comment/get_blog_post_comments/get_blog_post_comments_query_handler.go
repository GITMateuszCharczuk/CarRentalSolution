package queries

import (
	contract "blog-api/Application.contract/BlogPostComments/GetBlogPostComments"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_comment_repository"
	blog_repository "blog-api/Domain/repository_interfaces/blog_post_repository"
	"blog-api/Domain/responses"
	"context"
)

type GetBlogPostCommentsQueryHandler struct {
	blogPostCommentQueryRepository repository_interfaces.BlogPostCommentQueryRepository
	blogPostQueryRepository        blog_repository.BlogPostQueryRepository
}

func NewGetBlogPostCommentsQueryHandler(
	blogPostCommentQueryRepository repository_interfaces.BlogPostCommentQueryRepository,
	blogPostQueryRepository blog_repository.BlogPostQueryRepository,
) *GetBlogPostCommentsQueryHandler {
	return &GetBlogPostCommentsQueryHandler{
		blogPostCommentQueryRepository: blogPostCommentQueryRepository,
		blogPostQueryRepository:        blogPostQueryRepository,
	}
}

func (h *GetBlogPostCommentsQueryHandler) Handle(ctx context.Context, query *GetBlogPostCommentsQuery) (*contract.GetBlogPostCommentsResponse, error) {
	result, err := h.blogPostCommentQueryRepository.GetComments(
		query.BlogPostIds,
		query.UserIds,
		query.DateTimeFrom,
		query.DateTimeTo,
		&query.Pagination,
		&query.Sortable,
	)

	if err != nil {
		response := responses.NewResponse[contract.GetBlogPostCommentsResponse](500, "Failed to retrieve comments")
		return &response, nil
	}

	return &contract.GetBlogPostCommentsResponse{
		BaseResponse:    responses.NewBaseResponse(200, "Comments retrieved successfully"),
		PaginatedResult: *result,
	}, nil
}
