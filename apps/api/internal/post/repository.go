package post

import (
	"fmt"

	"gorm.io/gorm"
)

// Repository handles data persistence for posts
type Repository struct {
	posts map[string]*Post
	db    *gorm.DB
}

// NewRepository creates a new post repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		posts: make(map[string]*Post),
		db:    db,
	}
}

// Create adds a new post
func (r *Repository) Create(post *Post) (*Post, error) {
	r.posts[post.ID] = post
	return post, nil
}

// FindByID retrieves a post by ID from the database
func (r *Repository) FindByID(id string) (*Post, error) {
	var post Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("post not found")
		}
		return nil, err
	}
	return &post, nil
}

// FindAll retrieves all posts
func (r *Repository) FindAll() ([]*Post, error) {
	posts := make([]*Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

// Update updates a post
func (r *Repository) Update(post *Post) error {
	r.posts[post.ID] = post
	return nil
}

// Delete removes a post
func (r *Repository) Delete(id string) error {
	delete(r.posts, id)
	return nil
}
