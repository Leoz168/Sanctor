package comment

import "github.com/google/uuid"

// Repository defines persistence operations for comments.
type Repository interface {
	Create(comment *Comment) error
	FindByID(id uuid.UUID) (*Comment, error)
	FindByPostID(postID uuid.UUID) []*Comment
	Update(comment *Comment) error
	Delete(id uuid.UUID) error
	ExistsPost(postID uuid.UUID) bool
	ExistsUser(userID uuid.UUID) bool
}
