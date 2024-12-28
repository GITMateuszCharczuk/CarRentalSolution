package contract

type GetBlogPostCommentsCountRequest struct {
	BlogPostId string `json:"blog_post_id" validate:"required"`
}
