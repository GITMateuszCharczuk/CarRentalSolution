package repository_interfaces

type BlogPostLikeQueryRepository interface {
	GetLikesCount(blogPostID string, userID string) (int64, error)
}
