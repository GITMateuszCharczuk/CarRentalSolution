package commands

import (
	"context"
	contract "identity-api/Application.contract/BlogPostComments/DeketeBlogPostComment"
	"identity-api/Application/services"
	repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	"identity-api/Domain/responses"
	data_fetcher "identity-api/Domain/service_interfaces"
)

type DeleteBlogPostCommentCommandHandler struct {
	blogPostCommentCommandRepository repository_interfaces.BlogPostCommentCommandRepository
	blogPostCommentQueryRepository   repository_interfaces.BlogPostCommentQueryRepository
	dataFetcher                      data_fetcher.DataFetcher
}

func NewDeleteBlogPostCommentCommandHandler(
	blogPostCommentCommandRepository repository_interfaces.BlogPostCommentCommandRepository,
	blogPostCommentQueryRepository repository_interfaces.BlogPostCommentQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) *DeleteBlogPostCommentCommandHandler {
	return &DeleteBlogPostCommentCommandHandler{
		blogPostCommentCommandRepository: blogPostCommentCommandRepository,
		blogPostCommentQueryRepository:   blogPostCommentQueryRepository,
		dataFetcher:                      dataFetcher,
	}
}

func (h *DeleteBlogPostCommentCommandHandler) Handle(ctx context.Context, command *DeleteBlogPostCommentCommand) (*contract.DeleteBlogPostCommentResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.DeleteBlogPostCommentResponse](401, "Unauthorized")
		return &response, nil
	}

	authorId, err := h.blogPostCommentQueryRepository.GetCommentAuthorId(command.BlogPostCommentId)
	if err != nil || authorId == nil {
		response := responses.NewResponse[contract.DeleteBlogPostCommentResponse](404, "Comment not found")
		return &response, nil
	}

	if *authorId != userInfo.ID && !services.IsAdminOrSuperAdmin(userInfo.Roles) {
		response := responses.NewResponse[contract.DeleteBlogPostCommentResponse](403, "Not authorized to delete this comment")
		return &response, nil
	}

	err = h.blogPostCommentCommandRepository.RemoveComment(ctx, command.BlogPostCommentId)
	if err != nil {
		response := responses.NewResponse[contract.DeleteBlogPostCommentResponse](500, "Failed to delete comment")
		return &response, nil
	}

	return &contract.DeleteBlogPostCommentResponse{
		BaseResponse: responses.NewBaseResponse(200, "Comment deleted successfully"),
	}, nil
}
