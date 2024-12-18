package queries

import (
	contract "blog-api/Application.contract/BlogPosts/GetBlogPost"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_repository"
	"blog-api/Domain/responses"
	"context"
)

type GetBlogPostQueryHandler struct {
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository
}

func NewGetBlogPostQueryHandler(
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository,
) *GetBlogPostQueryHandler {
	return &GetBlogPostQueryHandler{
		blogPostQueryRepository: blogPostQueryRepository,
	}
}

func (h *GetBlogPostQueryHandler) Handle(ctx context.Context, query *GetBlogPostQuery) (*contract.GetBlogPostResponse, error) {
	blogPost, err := h.blogPostQueryRepository.GetBlogPostByID(query.ID)
	if err != nil {
		response := responses.NewResponse[contract.GetBlogPostResponse](500, "Failed to retrieve blog post")
		return &response, nil
	}

	if blogPost == nil {
		response := responses.NewResponse[contract.GetBlogPostResponse](404, "Blog post not found")
		return &response, nil
	}

	return &contract.GetBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Blog post retrieved successfully"),
		BlogPost:     *blogPost,
	}, nil
}
