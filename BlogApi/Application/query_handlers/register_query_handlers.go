package queries

import (
	get_blog_post "identity-api/Application/query_handlers/blog_post/get_blog_post"
	get_blog_posts "identity-api/Application/query_handlers/blog_post/get_blog_posts"
	get_blog_post_comments "identity-api/Application/query_handlers/blog_post_comment/get_blog_post_comments"
	get_likes_for_blog_post "identity-api/Application/query_handlers/blog_post_like/get_likes_for_blog_post"
	get_tags "identity-api/Application/query_handlers/blog_post_tag/get_tags"
	blog_post_comment_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	blog_post_like_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	blog_post_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	blog_post_tag_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_tag_repository"
	data_fetcher "identity-api/Domain/service_interfaces"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerGetBlogPostsQueryHandler(
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	handler := get_blog_posts.NewGetBlogPostsQueryHandler(blogPostQueryRepository, dataFetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetBlogPostQueryHandler(
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
) {
	handler := get_blog_post.NewGetBlogPostQueryHandler(blogPostQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetBlogPostCommentsQueryHandler(
	blogPostCommentQueryRepository blog_post_comment_repository_interfaces.BlogPostCommentQueryRepository,
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
) {
	handler := get_blog_post_comments.NewGetBlogPostCommentsQueryHandler(blogPostCommentQueryRepository, blogPostQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetLikesForBlogPostQueryHandler(
	blogPostLikeQueryRepository blog_post_like_repository_interfaces.BlogPostLikeQueryRepository,
) {
	handler := get_likes_for_blog_post.NewGetLikesForBlogPostQueryHandler(blogPostLikeQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetTagsQueryHandler(
	blogPostTagQueryRepository blog_post_tag_repository_interfaces.BlogPostTagQueryRepository,
) {
	handler := get_tags.NewGetTagsQueryHandler(blogPostTagQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterQueryHandlers(
	blogPostQueryRepository blog_post_repository_interfaces.BlogPostQueryRepository,
	blogPostCommentQueryRepository blog_post_comment_repository_interfaces.BlogPostCommentQueryRepository,
	blogPostLikeQueryRepository blog_post_like_repository_interfaces.BlogPostLikeQueryRepository,
	blogPostTagQueryRepository blog_post_tag_repository_interfaces.BlogPostTagQueryRepository,
	dataFetcher data_fetcher.DataFetcher,
) {
	registerGetBlogPostsQueryHandler(blogPostQueryRepository, dataFetcher)
	registerGetBlogPostQueryHandler(blogPostQueryRepository)
	registerGetBlogPostCommentsQueryHandler(blogPostCommentQueryRepository, blogPostQueryRepository)
	registerGetLikesForBlogPostQueryHandler(blogPostLikeQueryRepository)
	registerGetTagsQueryHandler(blogPostTagQueryRepository)
}
