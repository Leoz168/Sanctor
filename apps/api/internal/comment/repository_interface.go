package comment

// Repository defines persistence operations for comments.
type Repository interface {
	Create(comment *Comment) error
	FindByID(id string) (*Comment, error)
	FindByPostID(postID string) []*Comment
	Update(comment *Comment) error
	Delete(id string) error
	ExistsPost(postID string) bool
	ExistsUser(userID string) bool
}
