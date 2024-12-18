package commands

import (
	contract "blog-api/Application.contract/BlogPostComments/CreateBlogPostComment"
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_comment_repository"
	blog_repository "blog-api/Domain/repository_interfaces/blog_post_repository"
	"blog-api/Domain/responses"
	data_fetcher "blog-api/Domain/service_interfaces"
	"context"
)

type CreateBlogPostCommentCommandHandler struct {
	blogPostCommentCommandRepository repository_interfaces.BlogPostCommentCommandRepository
	blogPostQueryRepository          blog_repository.BlogPostQueryRepository
	dataFetcher                      data_fetcher.DataFetcher
}

func NewCreateBlogPostCommentCommandHandler(
	blogPostCommentCommandRepository repository_interfaces.BlogPostCommentCommandRepository,
	blogPostQueryRepository blog_repository.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) *CreateBlogPostCommentCommandHandler {
	return &CreateBlogPostCommentCommandHandler{
		blogPostCommentCommandRepository: blogPostCommentCommandRepository,
		blogPostQueryRepository:          blogPostQueryRepository,
		dataFetcher:                      dataFetcher,
	}
}

func (h *CreateBlogPostCommentCommandHandler) Handle(ctx context.Context, command *CreateBlogPostCommentCommand) (*contract.CreateBlogPostCommentResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.CreateBlogPostCommentResponse](401, "Unauthorized")
		return &response, nil
	}

	blogPost, err := h.blogPostQueryRepository.GetBlogPostByID(command.BlogPostId)
	if err != nil || blogPost == nil {
		response := responses.NewResponse[contract.CreateBlogPostCommentResponse](404, "Blog post not found")
		return &response, nil
	}

	comment := &models.BlogPostCommentModel{
		Description: command.Description,
		BlogPostId:  command.BlogPostId,
		UserId:      userInfo.ID,
	}

	commentID, err := h.blogPostCommentCommandRepository.AddComment(ctx, comment)
	if err != nil {
		response := responses.NewResponse[contract.CreateBlogPostCommentResponse](500, "Failed to create comment")
		return &response, nil
	}

	return &contract.CreateBlogPostCommentResponse{
		BaseResponse: responses.NewBaseResponse(200, "Comment created successfully"),
		Id:           commentID,
	}, nil
}
