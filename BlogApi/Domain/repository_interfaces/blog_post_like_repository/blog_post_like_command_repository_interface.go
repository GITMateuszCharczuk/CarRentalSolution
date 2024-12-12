package repository_interfaces

type BlogPostLikeCommandRepository interface {
	AddLike(blogPostID string, userID string) error
	RemoveLike(blogPostID string, userID string) error
}
