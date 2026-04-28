package post

import "github.com/google/uuid"

// RepositoryInterface defines the contract for post data persistence
type RepositoryInterface interface {
	CreateWithLinks(post *Post, groupIDs []uuid.UUID, institutionIDs []uuid.UUID) (*Post, error)
	FindByID(id uuid.UUID) (*Post, error)
	FindAll() ([]*Post, error)
	Search(filters PostSearchFilters) ([]*Post, error)
	Update(post *Post) error
	Delete(id uuid.UUID) error
}
