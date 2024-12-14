package queries

import (
	"context"
	contract "identity-api/Application.contract/BlogPosts/GetBlogPost"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	"identity-api/Domain/responses"
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
		BaseResponse:     responses.NewBaseResponse(200, "Blog post retrieved successfully"),
		Id:               blogPost.Id,
		Heading:          blogPost.Heading,
		PageTitle:        blogPost.PageTitle,
		Content:          blogPost.Content,
		ShortDescription: blogPost.ShortDescription,
		FeaturedImageUrl: blogPost.FeaturedImageUrl,
		UrlHandle:        blogPost.UrlHandle,
		PublishedDate:    blogPost.PublishedDate,
		Author:           blogPost.AuthorName,
	}, nil
}
