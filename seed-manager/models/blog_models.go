package models

import "time"

type BlogPostModel struct {
	ID        string
	Title     string
	Content   string
	AuthorID  string
	ImageId   string
	CreatedAt time.Time
}

type BlogCommentModel struct {
	ID        string
	Content   string
	AuthorID  string
	PostID    string
	CreatedAt time.Time
}

type BlogTagModel struct {
	ID   string
	Name string
}
