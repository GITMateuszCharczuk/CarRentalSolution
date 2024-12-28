package queries

import (
	contract "blog-api/Application.contract/BlogPostComments/GetBlogPostCommentsCount"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_comment_repository"
	"blog-api/Domain/responses"
	"context"
)

type GetBlogPostCommentsCountQueryHandler struct {
	blogPostCommentQueryRepository repository_interfaces.BlogPostCommentQueryRepository
}

func NewGetBlogPostCommentsCountQueryHandler(
	blogPostCommentQueryRepository repository_interfaces.BlogPostCommentQueryRepository,
) *GetBlogPostCommentsCountQueryHandler {
	return &GetBlogPostCommentsCountQueryHandler{
		blogPostCommentQueryRepository: blogPostCommentQueryRepository,
	}
}

func (h *GetBlogPostCommentsCountQueryHandler) Handle(ctx context.Context, query *GetBlogPostCommentsCountQuery) (*contract.GetBlogPostCommentsCountResponse, error) {
	result, err := h.blogPostCommentQueryRepository.GetCommentsCount(
		query.BlogPostId,
	)

	if err != nil {
		response := responses.NewResponse[contract.GetBlogPostCommentsCountResponse](500, "Failed to retrieve comments")
		return &response, nil
	}

	return &contract.GetBlogPostCommentsCountResponse{
		BaseResponse: responses.NewBaseResponse(200, "Comments count retrieved successfully"),
		Count:        result,
	}, nil
}
