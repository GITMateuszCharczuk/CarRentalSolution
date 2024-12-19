package commands

import (
	"context"
	contract "rental-api/Application.contract/BlogPosts/DeleteBlogPost"
	"rental-api/Application/services"
	"rental-api/Domain/constants"
	repository_interfaces "rental-api/Domain/repository_interfaces/blog_post_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type DeleteBlogPostCommandHandler struct {
	blogPostCommandRepository repository_interfaces.BlogPostCommandRepository
	blogPostQueryRepository   repository_interfaces.BlogPostQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewDeleteBlogPostCommandHandler(
	blogPostCommandRepository repository_interfaces.BlogPostCommandRepository,
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *DeleteBlogPostCommandHandler {
	return &DeleteBlogPostCommandHandler{
		blogPostCommandRepository: blogPostCommandRepository,
		blogPostQueryRepository:   blogPostQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *DeleteBlogPostCommandHandler) Handle(ctx context.Context, command *DeleteBlogPostCommand) (*contract.DeleteBlogPostResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil || !services.IsAdminOrSuperAdmin(userInfo.Roles) {
		response := responses.NewResponse[contract.DeleteBlogPostResponse](401, "Unauthorized")
		return &response, nil
	}

	authorId, err := h.blogPostQueryRepository.GetBlogPostAuthorId(command.ID)
	if err != nil || authorId == nil {
		response := responses.NewResponse[contract.DeleteBlogPostResponse](404, "Blog post not found")
		return &response, nil
	}

	if !services.IsRole(constants.SuperAdmin, userInfo.Roles) && *authorId != userInfo.ID {
		response := responses.NewResponse[contract.DeleteBlogPostResponse](403, "Not authorized to delete this blog post")
		return &response, nil
	}

	err = h.blogPostCommandRepository.DeleteBlogPost(ctx, command.ID)
	if err != nil {
		response := responses.NewResponse[contract.DeleteBlogPostResponse](500, "Failed to delete blog post")
		return &response, nil
	}

	return &contract.DeleteBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Blog post deleted successfully"),
	}, nil
}
