package post

// RepositoryInterface defines the contract for post data persistence
type RepositoryInterface interface {
	CreateWithLinks(post *Post, groupIDs []string, institutionIDs []string) (*Post, error)
	FindByID(id string) (*Post, error)
	FindAll() ([]*Post, error)
	Update(post *Post) error
	Delete(id string) error
}
