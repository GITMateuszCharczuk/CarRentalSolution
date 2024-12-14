package commands

import (
	"context"
	contract "identity-api/Application.contract/BlogPostLikes/CreateLikeForBlogPost"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	blog_repository "identity-api/Domain/repository_interfaces/blog_post_repository"
	"identity-api/Domain/responses"
	data_fetcher "identity-api/Domain/service_interfaces"
)

type CreateLikeForBlogPostCommandHandler struct {
	blogPostLikeCommandRepository repository_interfaces.BlogPostLikeCommandRepository
	blogPostQueryRepository       blog_repository.BlogPostQueryRepository
	dataFetcher                   data_fetcher.DataFetcher
}

func NewCreateLikeForBlogPostCommandHandler(
	blogPostLikeCommandRepository repository_interfaces.BlogPostLikeCommandRepository,
	blogPostQueryRepository blog_repository.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) *CreateLikeForBlogPostCommandHandler {
	return &CreateLikeForBlogPostCommandHandler{
		blogPostLikeCommandRepository: blogPostLikeCommandRepository,
		blogPostQueryRepository:       blogPostQueryRepository,
		dataFetcher:                   dataFetcher,
	}
}

func (h *CreateLikeForBlogPostCommandHandler) Handle(ctx context.Context, command *CreateLikeForBlogPostCommand) (*contract.CreateLikeForBlogPostResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.CreateLikeForBlogPostResponse](401, "Unauthorized")
		return &response, nil
	}

	blogPost, err := h.blogPostQueryRepository.GetBlogPostByID(command.BlogPostId)
	if err != nil || blogPost == nil {
		response := responses.NewResponse[contract.CreateLikeForBlogPostResponse](404, "Blog post not found")
		return &response, nil
	}

	err = h.blogPostLikeCommandRepository.AddLike(ctx, command.BlogPostId, userInfo.ID)
	if err != nil {
		response := responses.NewResponse[contract.CreateLikeForBlogPostResponse](500, "Failed to add like")
		return &response, nil
	}

	return &contract.CreateLikeForBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Like added successfully"),
	}, nil
}
