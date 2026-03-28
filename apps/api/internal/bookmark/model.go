package bookmark

import (
	"time"

	"sanctor/internal/post"
)

// Bookmark represents a user's saved post.
type Bookmark struct {
	UserID    string    `json:"userId" gorm:"type:uuid;primaryKey;not null;index:idx_post_bookmarks_user_id"`
	PostID    string    `json:"postId" gorm:"type:uuid;primaryKey;not null;index:idx_post_bookmarks_post_id"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

// TableName ensures a stable junction table name.
func (Bookmark) TableName() string {
	return "post_bookmarks"
}

// CreateBookmarkRequest is the payload for creating a bookmark.
type CreateBookmarkRequest struct {
	UserID string `json:"userId"`
	PostID string `json:"postId"`
}

// BookmarkStatusResponse returns bookmark state for one user+post pair.
type BookmarkStatusResponse struct {
	UserID       string `json:"userId"`
	PostID       string `json:"postId"`
	IsBookmarked bool   `json:"isBookmarked"`
}

// BookmarkedPost is a full post row returned by bookmark listing APIs.
type BookmarkedPost = post.Post
