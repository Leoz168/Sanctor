package user

import "errors"

// Repository handles data persistence for users
type Repository struct {
	users map[string]*User
}

// NewRepository creates a new user repository
func NewRepository() *Repository {
	return &Repository{
		users: make(map[string]*User),
	}
}

// Create adds a new user to the repository
func (r *Repository) Create(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	r.users[user.ID] = user
	return nil
}

// FindByID retrieves a user by ID
func (r *Repository) FindByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// FindAll retrieves all users
func (r *Repository) FindAll() []*User {
	userList := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList
}

// Update updates an existing user
func (r *Repository) Update(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

// Delete removes a user from the repository
func (r *Repository) Delete(id string) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

// ExistsByEmail checks if a user with the given email exists
func (r *Repository) ExistsByEmail(email string) bool {
	for _, user := range r.users {
		if user.Email == email {
			return true
		}
	}
	return false
}

// ExistsByUsername checks if a user with the given username exists
func (r *Repository) ExistsByUsername(username string) bool {
	for _, user := range r.users {
		if user.Username == username {
			return true
		}
	}
	return false
}

// FindByEmail retrieves a user by email
func (r *Repository) FindByEmail(email string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// FindByUsername retrieves a user by username
func (r *Repository) FindByUsername(username string) (*User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
