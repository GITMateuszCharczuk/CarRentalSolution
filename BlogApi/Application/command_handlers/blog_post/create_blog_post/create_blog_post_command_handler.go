package commands

import (
	contract "blog-api/Application.contract/BlogPosts/CreateBlogPost"
	models "blog-api/Domain/models/domestic"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_repository"
	"blog-api/Domain/responses"
	data_fetcher "blog-api/Domain/service_interfaces"
	"context"
	"time"
)

type CreateBlogPostCommandHandler struct {
	blogPostCommandRepository repository_interfaces.BlogPostCommandRepository
	dataFetcher               data_fetcher.DataFetcher
}

func NewCreateBlogPostCommandHandler(
	blogPostCommandRepository repository_interfaces.BlogPostCommandRepository,
	dataFetcher data_fetcher.DataFetcher,
) *CreateBlogPostCommandHandler {
	return &CreateBlogPostCommandHandler{
		blogPostCommandRepository: blogPostCommandRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *CreateBlogPostCommandHandler) Handle(ctx context.Context, command *CreateBlogPostCommand) (*contract.CreateBlogPostResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.CreateBlogPostResponse](401, "Unauthorized")
		return &response, nil
	}

	now := time.Now()
	blogPost := &models.BlogPostRequestModel{
		Heading:          command.Heading,
		PageTitle:        command.PageTitle,
		Content:          command.Content,
		ShortDescription: command.ShortDescription,
		FeaturedImageUrl: command.FeaturedImageUrl,
		UrlHandle:        command.UrlHandle,
		AuthorId:         userInfo.ID,
		AuthorName:       userInfo.Name,
		Visible:          command.Visible,
		CreatedAt:        now.Format(time.RFC3339),
		Tags:             command.Tags,
	}

	result, err := h.blogPostCommandRepository.CreateBlogPost(ctx, blogPost)
	if err != nil {
		response := responses.NewResponse[contract.CreateBlogPostResponse](500, "Failed to create blog post")
		return &response, nil
	}

	return &contract.CreateBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Blog post created successfully"),
		Id:           *result,
	}, nil
}
