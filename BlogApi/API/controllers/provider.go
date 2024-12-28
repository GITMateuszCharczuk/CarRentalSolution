package controllers

import (
	base "blog-api/API/controllers/base"
	blog "blog-api/API/controllers/blog_post"
	comment "blog-api/API/controllers/blog_post_comment"
	like "blog-api/API/controllers/blog_post_like"
	tag "blog-api/API/controllers/blog_post_tag"

	"github.com/google/wire"
)

var ControllerSet = wire.NewSet(
	NewControllers,
	ProvideControllers,
	blog.NewCreateBlogPostController,
	blog.NewGetBlogPostController,
	blog.NewGetBlogPostsController,
	blog.NewUpdateBlogPostController,
	blog.NewDeleteBlogPostController,
	comment.NewCreateBlogPostCommentController,
	comment.NewDeleteBlogPostCommentController,
	comment.NewGetBlogPostCommentsController,
	comment.NewGetBlogPostCommentsCountController,
	like.NewCreateLikeForBlogPostController,
	like.NewDeleteLikeForBlogPostController,
	like.NewGetLikesForBlogPostController,
	tag.NewGetTagsController,
)

type Controllers struct {
	All []base.Controller
}

func NewControllers(all []base.Controller) *Controllers {
	return &Controllers{All: all}
}

func ProvideControllers(
	// Blog Post Controllers
	createBlogPostController *blog.CreateBlogPostController,
	getBlogPostController *blog.GetBlogPostController,
	getBlogPostsController *blog.GetBlogPostsController,
	updateBlogPostController *blog.UpdateBlogPostController,
	deleteBlogPostController *blog.DeleteBlogPostController,

	// Blog Post Comment Controllers
	createBlogPostCommentController *comment.CreateBlogPostCommentController,
	deleteBlogPostCommentController *comment.DeleteBlogPostCommentController,
	getBlogPostCommentsController *comment.GetBlogPostCommentsController,
	getBlogPostCommentsCountController *comment.GetBlogPostCommentsCountController,

	// Blog Post Like Controllers
	createLikeForBlogPostController *like.CreateLikeForBlogPostController,
	deleteLikeForBlogPostController *like.DeleteLikeForBlogPostController,
	getLikesForBlogPostController *like.GetLikesForBlogPostController,

	// Blog Post Tag Controllers
	getTagsController *tag.GetTagsController,
) []base.Controller {
	return []base.Controller{
		// Blog Post Controllers
		createBlogPostController,
		getBlogPostController,
		getBlogPostsController,
		updateBlogPostController,
		deleteBlogPostController,

		// Blog Post Comment Controllers
		createBlogPostCommentController,
		deleteBlogPostCommentController,
		getBlogPostCommentsController,
		getBlogPostCommentsCountController,

		// Blog Post Like Controllers
		createLikeForBlogPostController,
		deleteLikeForBlogPostController,
		getLikesForBlogPostController,

		// Blog Post Tag Controllers
		getTagsController,
	}
}

var WireSet = wire.NewSet(
	ProvideControllers,
	NewControllers,
	blog.NewCreateBlogPostController,
	blog.NewGetBlogPostController,
	blog.NewGetBlogPostsController,
	blog.NewUpdateBlogPostController,
	blog.NewDeleteBlogPostController,
	comment.NewCreateBlogPostCommentController,
	comment.NewDeleteBlogPostCommentController,
	comment.NewGetBlogPostCommentsController,
	comment.NewGetBlogPostCommentsCountController,
	like.NewCreateLikeForBlogPostController,
	like.NewDeleteLikeForBlogPostController,
	like.NewGetLikesForBlogPostController,
	tag.NewGetTagsController,
)
