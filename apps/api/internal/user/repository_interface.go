package user

import "github.com/google/uuid"

// Repository defines the interface for user data access
type Repository interface {
	Create(user *User) error
	FindByID(id uuid.UUID) (*User, error)
	FindAll() []*User
	Update(user *User) error
	Delete(id uuid.UUID) error
	ExistsByEmail(email string) bool
	ExistsByUsername(username string) bool
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
}
