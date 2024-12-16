package queries

import (
	"context"
	contract "identity-api/Application.contract/BlogPosts/GetBlogPosts"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	"identity-api/Domain/responses"
	data_fetcher "identity-api/Domain/service_interfaces"
)

type GetBlogPostsQueryHandler struct {
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository
	dataFetcher             data_fetcher.DataFetcher
}

func NewGetBlogPostsQueryHandler(
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) *GetBlogPostsQueryHandler {
	return &GetBlogPostsQueryHandler{
		blogPostQueryRepository: blogPostQueryRepository,
		dataFetcher:             dataFetcher,
	}
}

func (h *GetBlogPostsQueryHandler) Handle(ctx context.Context, query *GetBlogPostsQuery) (*contract.GetBlogPostsResponse, error) {
	// tokenInfo, err := h.dataFetcher.ValidateToken(query.JwtToken)
	// if err != nil || !services.IsAdminOrSuperAdmin(tokenInfo.Roles) {
	// 	query.Visible = true
	// } // TODO

	result, err := h.blogPostQueryRepository.GetBlogPosts(
		&query.Pagination,
		&query.Sortable,
		query.Ids,
		query.DateTimeFrom,
		query.DateTimeTo,
		query.AuthorIds,
		query.Tags,
		query.Visible,
	)

	if err != nil {
		response := responses.NewResponse[contract.GetBlogPostsResponse](500, "Failed to retrieve blog posts")
		return &response, nil
	}

	return &contract.GetBlogPostsResponse{
		BaseResponse:    responses.NewBaseResponse(200, "Blog posts retrieved successfully"),
		PaginatedResult: *result,
	}, nil
}
