package commands

import (
	create_blog_post "blog-api/Application/command_handlers/blog_post/create_blog_post"
	delete_blog_post "blog-api/Application/command_handlers/blog_post/delete_blog_post"
	update_blog_post "blog-api/Application/command_handlers/blog_post/update_blog_post"
	create_blog_post_comment "blog-api/Application/command_handlers/blog_post_comment/create_blog_post_comment"
	delete_blog_post_comment "blog-api/Application/command_handlers/blog_post_comment/delete_blog_post_comment"
	create_like_for_blog_post "blog-api/Application/command_handlers/blog_post_like/create_like_for_blog_post"
	delete_like_for_blog_post "blog-api/Application/command_handlers/blog_post_like/delete_like_for_blog_post"
	blog_post_comment_repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_comment_repository"
	blog_post_like_repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_like_repository"
	blog_post_repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_repository"
	data_fetcher "blog-api/Domain/service_interfaces"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerCreateBlogPostCommandHandler(
	blogCommandRepository blog_post_repository_interfaces.BlogPostCommandRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := create_blog_post.NewCreateBlogPostCommandHandler(blogCommandRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerUpdateBlogPostCommandHandler(
	blogCommandRepository blog_post_repository_interfaces.BlogPostCommandRepository,
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := update_blog_post.NewUpdateBlogPostCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteBlogPostCommandHandler(
	blogCommandRepository blog_post_repository_interfaces.BlogPostCommandRepository,
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := delete_blog_post.NewDeleteBlogPostCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerCreateBlogPostCommentCommandHandler(
	blogCommandRepository blog_post_comment_repository_interfaces.BlogPostCommentCommandRepository,
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := create_blog_post_comment.NewCreateBlogPostCommentCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteBlogPostCommentCommandHandler(
	blogCommandRepository blog_post_comment_repository_interfaces.BlogPostCommentCommandRepository,
	blogPostQueryRepository blog_post_comment_repository_interfaces.BlogPostCommentQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := delete_blog_post_comment.NewDeleteBlogPostCommentCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerCreateLikeForBlogPostCommandHandler(
	blogCommandRepository blog_post_like_repository_interfaces.BlogPostLikeCommandRepository,
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := create_like_for_blog_post.NewCreateLikeForBlogPostCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteLikeForBlogPostCommandHandler(
	blogCommandRepository blog_post_like_repository_interfaces.BlogPostLikeCommandRepository,
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := delete_like_for_blog_post.NewDeleteLikeForBlogPostCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterCommandHandlers(
	blogCommandRepository blog_post_repository_interfaces.BlogPostCommandRepository,
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	blogPostCommentCommandRepository blog_post_comment_repository_interfaces.BlogPostCommentCommandRepository,
	blogPostCommentQueryRepository blog_post_comment_repository_interfaces.BlogPostCommentQueryRepository,
	blogPostLikeCommandRepository blog_post_like_repository_interfaces.BlogPostLikeCommandRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	registerCreateBlogPostCommandHandler(blogCommandRepository, dataFetcher)
	registerUpdateBlogPostCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	registerDeleteBlogPostCommandHandler(blogCommandRepository, blogPostQueryRepository, dataFetcher)
	registerCreateBlogPostCommentCommandHandler(blogPostCommentCommandRepository, blogPostQueryRepository, dataFetcher)
	registerDeleteBlogPostCommentCommandHandler(blogPostCommentCommandRepository, blogPostCommentQueryRepository, dataFetcher)
	registerCreateLikeForBlogPostCommandHandler(blogPostLikeCommandRepository, blogPostQueryRepository, dataFetcher)
	registerDeleteLikeForBlogPostCommandHandler(blogPostLikeCommandRepository, blogPostQueryRepository, dataFetcher)
}
