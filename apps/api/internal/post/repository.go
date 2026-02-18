package post

// Repository handles data persistence for posts
type Repository struct {
	posts map[string]*Post
}

// NewRepository creates a new post repository
func NewRepository() *Repository {
	return &Repository{
		posts: make(map[string]*Post),
	}
}

// Create adds a new post
func (r *Repository) Create(post *Post) (*Post, error) {
	r.posts[post.ID] = post
	return post, nil
}

// FindByID retrieves a post by ID
func (r *Repository) FindByID(id string) (*Post, error) {
	if post, ok := r.posts[id]; ok {
		return post, nil
	}
	return nil, nil
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
