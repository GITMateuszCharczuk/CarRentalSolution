package entities

import (
	"time"

	"github.com/google/uuid"
)

type BlogPostTagEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"unique" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Navigation Properties
	BlogPosts []BlogPostEntity `gorm:"many2many:blog_post_tags;" json:"blog_posts"`
}
