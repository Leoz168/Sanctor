package comment

import "time"

// Comment represents a comment under a post.
type Comment struct {
	ID              string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PostID          string     `json:"postId" gorm:"type:uuid;not null;index:idx_comments_post_id"`
	CreatedByUserID string     `json:"createdByUserId" gorm:"type:uuid;not null;index:idx_comments_created_by_user_id"`
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
