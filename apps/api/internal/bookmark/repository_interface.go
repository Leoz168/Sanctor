package bookmark

import "github.com/google/uuid"

// Repository defines persistence operations for bookmarks.
type Repository interface {
	Create(bookmark *Bookmark) error
	Delete(userID, postID uuid.UUID) error
	FindByUserID(userID uuid.UUID) ([]*Bookmark, error)
	FindPostsByUserID(userID uuid.UUID) ([]*BookmarkedPost, error)
	Exists(userID, postID uuid.UUID) (bool, error)
	ExistsPost(postID uuid.UUID) bool
	ExistsUser(userID uuid.UUID) bool
}
