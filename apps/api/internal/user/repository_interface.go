package user

// Repository defines the interface for user data access
type Repository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindAll() []*User
	Update(user *User) error
	Delete(id string) error
	ExistsByEmail(email string) bool
	ExistsByUsername(username string) bool
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
}
