package entities

import (
	"time"

	"github.com/google/uuid"
)

type BlogPostEntity struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Heading          string    `json:"heading"`
	PageTitle        string    `json:"page_title"`
	Content          string    `json:"content"`
	ShortDescription string    `json:"short_description"`
	FeaturedImageUrl string    `json:"featured_image_url"`
	UrlHandle        string    `json:"url_handle"`
	PublishedDate    time.Time `json:"published_date"`
	Author           string    `json:"author"`
	UserId           uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	Visible          bool      `json:"visible"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	// Navigation Properties
	Tags     []BlogPostTagEntity     `gorm:"many2many:blog_post_tags;constraint:OnDelete:CASCADE;" json:"tags"`
	Likes    []BlogPostLikeEntity    `gorm:"foreignKey:BlogPostID;constraint:OnDelete:CASCADE;" json:"likes"`
	Comments []BlogPostCommentEntity `gorm:"foreignKey:BlogPostID;constraint:OnDelete:CASCADE;" json:"comments"`
}
