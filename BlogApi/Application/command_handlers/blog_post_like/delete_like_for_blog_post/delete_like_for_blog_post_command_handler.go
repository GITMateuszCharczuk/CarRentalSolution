package commands

import (
	"context"
	contract "identity-api/Application.contract/BlogPostLikes/DeleteLikeForBlogPost"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	blog_repository "identity-api/Domain/repository_interfaces/blog_post_repository"
	"identity-api/Domain/responses"
	data_fetcher "identity-api/Domain/service_interfaces"
)

type DeleteLikeForBlogPostCommandHandler struct {
	blogPostLikeCommandRepository repository_interfaces.BlogPostLikeCommandRepository
	blogPostQueryRepository       blog_repository.BlogPostQueryRepository
	dataFetcher                   data_fetcher.DataFetcher
}

func NewDeleteLikeForBlogPostCommandHandler(
	blogPostLikeCommandRepository repository_interfaces.BlogPostLikeCommandRepository,
	blogPostQueryRepository blog_repository.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) *DeleteLikeForBlogPostCommandHandler {
	return &DeleteLikeForBlogPostCommandHandler{
		blogPostLikeCommandRepository: blogPostLikeCommandRepository,
		blogPostQueryRepository:       blogPostQueryRepository,
		dataFetcher:                   dataFetcher,
	}
}

func (h *DeleteLikeForBlogPostCommandHandler) Handle(ctx context.Context, command *DeleteLikeForBlogPostCommand) (*contract.DeleteLikeForBlogPostResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.DeleteLikeForBlogPostResponse](401, "Unauthorized")
		return &response, nil
	}

	blogPost, err := h.blogPostQueryRepository.GetBlogPostByID(command.BlogPostId)
	if err != nil || blogPost == nil {
		response := responses.NewResponse[contract.DeleteLikeForBlogPostResponse](404, "Blog post not found")
		return &response, nil
	}

	err = h.blogPostLikeCommandRepository.RemoveLike(ctx, command.BlogPostId, userInfo.ID)
	if err != nil {
		response := responses.NewResponse[contract.DeleteLikeForBlogPostResponse](500, "Failed to remove like")
		return &response, nil
	}

	return &contract.DeleteLikeForBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Like removed successfully"),
	}, nil
}
