package queries

import (
	"context"
	contract "rental-api/Application.contract/Tags/GetTags"
	repository_interfaces "rental-api/Domain/repository_interfaces/blog_post_tag_repository"
	"rental-api/Domain/responses"
)

type GetTagsQueryHandler struct {
	blogPostTagQueryRepository repository_interfaces.BlogPostTagQueryRepository
}

func NewGetTagsQueryHandler(
	blogPostTagQueryRepository repository_interfaces.BlogPostTagQueryRepository,
) *GetTagsQueryHandler {
	return &GetTagsQueryHandler{
		blogPostTagQueryRepository: blogPostTagQueryRepository,
	}
}

func (h *GetTagsQueryHandler) Handle(ctx context.Context, query *GetTagsQuery) (*contract.GetTagsResponse, error) {
	tags, err := h.blogPostTagQueryRepository.GetTagsByBlogPostID(query.BlogPostId, query.Sortable)
	if err != nil {
		response := responses.NewResponse[contract.GetTagsResponse](500, "Failed to retrieve tags")
		return &response, nil
	}

	return &contract.GetTagsResponse{
		BaseResponse: responses.NewBaseResponse(200, "Tags retrieved successfully"),
		Items:        *tags,
	}, nil
}
