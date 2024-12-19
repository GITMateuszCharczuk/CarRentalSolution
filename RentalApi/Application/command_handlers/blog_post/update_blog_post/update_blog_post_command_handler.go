package commands

import (
	"context"
	"log"
	contract "rental-api/Application.contract/BlogPosts/UpdateBlogPost"
	"rental-api/Application/services"
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/blog_post_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type UpdateBlogPostCommandHandler struct {
	blogPostCommandRepository repository_interfaces.BlogPostCommandRepository
	blogPostQueryRepository   repository_interfaces.BlogPostQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewUpdateBlogPostCommandHandler(
	blogPostCommandRepository repository_interfaces.BlogPostCommandRepository,
	blogPostQueryRepository repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *UpdateBlogPostCommandHandler {
	return &UpdateBlogPostCommandHandler{
		blogPostCommandRepository: blogPostCommandRepository,
		blogPostQueryRepository:   blogPostQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *UpdateBlogPostCommandHandler) Handle(ctx context.Context, command *UpdateBlogPostCommand) (*contract.UpdateBlogPostResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.UpdateBlogPostResponse](401, "Unauthorized")
		return &response, nil
	}

	existingPost, err := h.blogPostQueryRepository.GetBlogPostByID(command.Id)
	if err != nil || existingPost == nil {
		response := responses.NewResponse[contract.UpdateBlogPostResponse](404, "Blog post not found")
		return &response, nil
	}
	authorId, err := h.blogPostQueryRepository.GetBlogPostAuthorId(command.Id)
	if err != nil || authorId == nil {
		response := responses.NewResponse[contract.UpdateBlogPostResponse](404, "Blog post not found")
		return &response, nil
	}

	if *authorId != userInfo.ID || !services.IsAdminOrSuperAdmin(userInfo.Roles) {
		response := responses.NewResponse[contract.UpdateBlogPostResponse](403, "Not authorized to update this blog post")
		return &response, nil
	}

	blogPost := &models.BlogPostRequestModel{
		Id:               command.Id,
		Heading:          command.Heading,
		PageTitle:        command.PageTitle,
		Content:          command.Content,
		ShortDescription: command.ShortDescription,
		FeaturedImageUrl: command.FeaturedImageUrl,
		UrlHandle:        command.UrlHandle,
		Visible:          command.Visible,
		Tags:             command.Tags,
		CreatedAt:        existingPost.CreatedAt,
		AuthorId:         *authorId,
		AuthorName:       existingPost.AuthorName,
	}

	err = h.blogPostCommandRepository.UpdateBlogPost(ctx, blogPost)
	log.Println(err)
	if err != nil {
		response := responses.NewResponse[contract.UpdateBlogPostResponse](500, "Failed to update blog post")
		return &response, nil
	}

	return &contract.UpdateBlogPostResponse{
		BaseResponse: responses.NewBaseResponse(200, "Blog post updated successfully"),
	}, nil
}
