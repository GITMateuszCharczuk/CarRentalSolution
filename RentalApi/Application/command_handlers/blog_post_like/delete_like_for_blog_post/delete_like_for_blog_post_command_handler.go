package commands

import (
	"context"
	contract "rental-api/Application.contract/BlogPostLikes/DeleteLikeForBlogPost"
	repository_interfaces "rental-api/Domain/repository_interfaces/blog_post_like_repository"
	blog_repository "rental-api/Domain/repository_interfaces/blog_post_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type DeleteLikeForBlogPostCommandHandler struct {
	blogPostLikeCommandRepository repository_interfaces.BlogPostLikeCommandRepository
	blogPostQueryRepository       blog_repository.BlogPostQueryRepository
	dataFetcher                   data_fetcher.MicroserviceConnector
}

func NewDeleteLikeForBlogPostCommandHandler(
	blogPostLikeCommandRepository repository_interfaces.BlogPostLikeCommandRepository,
	blogPostQueryRepository blog_repository.BlogPostQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
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
		if err.Error() == "record not found" {
			response := responses.NewResponse[contract.DeleteLikeForBlogPostResponse](404, "Like not found")
			return &response, nil
		}
		response := responses.NewResponse[contract.DeleteLikeForBlogPostResponse](500, "Failed to remove like")
		return &response, nil
	}

	return &contract.DeleteLikeForBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Like removed successfully"),
	}, nil
}
