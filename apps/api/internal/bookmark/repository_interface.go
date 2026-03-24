package bookmark

// Repository defines persistence operations for bookmarks.
type Repository interface {
	Create(bookmark *Bookmark) error
	Delete(userID, postID string) error
	FindByUserID(userID string) ([]*Bookmark, error)
	FindPostsByUserID(userID string) ([]*BookmarkedPost, error)
	Exists(userID, postID string) (bool, error)
	ExistsPost(postID string) bool
	ExistsUser(userID string) bool
}
