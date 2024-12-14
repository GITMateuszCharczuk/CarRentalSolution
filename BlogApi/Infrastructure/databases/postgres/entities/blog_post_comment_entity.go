package entities

import (
	"time"

	"github.com/google/uuid"
)

type BlogPostCommentEntity struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BlogPostID  uuid.UUID `gorm:"type:uuid;index" json:"blog_post_id"`
	UserID      uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Navigation Properties
	BlogPost BlogPostEntity `gorm:"foreignKey:BlogPostID" json:"blog_post"`
}
