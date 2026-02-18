package user

import "errors"

// InMemoryRepository handles data persistence for users in memory
type InMemoryRepository struct {
	users map[string]*User
}

// NewRepository creates a new in-memory user repository
func NewRepository() Repository {
	return &InMemoryRepository{
		users: make(map[string]*User),
	}
}

// Create adds a new user to the repository
func (r *InMemoryRepository) Create(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	r.users[user.ID] = user
	return nil
}

// FindByID retrieves a user by ID
func (r *InMemoryRepository) FindByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// FindAll retrieves all users
func (r *InMemoryRepository) FindAll() []*User {
	userList := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList
}

// Update updates an existing user
func (r *InMemoryRepository) Update(user *User) error {
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
func (r *InMemoryRepository) Delete(id string) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

// ExistsByEmail checks if a user with the given email exists
func (r *InMemoryRepository) ExistsByEmail(email string) bool {
	for _, user := range r.users {
		if user.Email == email {
			return true
		}
	}
	return false
}

// ExistsByUsername checks if a user with the given username exists
func (r *InMemoryRepository) ExistsByUsername(username string) bool {
	for _, user := range r.users {
		if user.Username == username {
			return true
		}
	}
	return false
}

// FindByEmail retrieves a user by email
func (r *InMemoryRepository) FindByEmail(email string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// FindByUsername retrieves a user by username
func (r *InMemoryRepository) FindByUsername(username string) (*User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
