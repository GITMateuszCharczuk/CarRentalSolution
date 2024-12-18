package queries

import (
	contract "blog-api/Application.contract/BlogPostLikes/GetLikesForBlogPost"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_like_repository"
	"blog-api/Domain/responses"
	"context"
)

type GetLikesForBlogPostQueryHandler struct {
	blogPostLikeQueryRepository repository_interfaces.BlogPostLikeQueryRepository
}

func NewGetLikesForBlogPostQueryHandler(
	blogPostLikeQueryRepository repository_interfaces.BlogPostLikeQueryRepository,
) *GetLikesForBlogPostQueryHandler {
	return &GetLikesForBlogPostQueryHandler{
		blogPostLikeQueryRepository: blogPostLikeQueryRepository,
	}
}

func (h *GetLikesForBlogPostQueryHandler) Handle(ctx context.Context, query *GetLikesForBlogPostQuery) (*contract.GetLikesForBlogPostResponse, error) {
	count, err := h.blogPostLikeQueryRepository.GetLikesCount(query.BlogPostId, "")

	if err != nil {
		response := responses.NewResponse[contract.GetLikesForBlogPostResponse](500, "Failed to retrieve likes count")
		return &response, nil
	}

	return &contract.GetLikesForBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Likes count retrieved successfully"),
		TotalCount:   int(count),
	}, nil
}
