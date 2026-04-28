package comment

import (
	"time"

	"github.com/google/uuid"
)

// Comment represents a comment under a post.
type Comment struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PostID          uuid.UUID  `json:"postId" gorm:"type:uuid;not null;index:idx_comments_post_id"`
	CreatedByUserID uuid.UUID  `json:"createdByUserId" gorm:"type:uuid;not null;index:idx_comments_created_by_user_id"`
	Content         string     `json:"content" gorm:"type:text;not null"`
	CreatedAt       time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}

// CreateCommentRequest represents the request body for creating a comment.
type CreateCommentRequest struct {
	PostID          string `json:"postId"`
	CreatedByUserID string `json:"createdByUserId"`
	Content         string `json:"content"`
}

// UpdateCommentRequest represents the request body for updating a comment.
type UpdateCommentRequest struct {
	Content string `json:"content"`
}
