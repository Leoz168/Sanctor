package user

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Name         string     `json:"name"`
	Email        string     `json:"email" gorm:"unique"`
	Username     string     `json:"username"`
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	PasswordHash string     `json:"-"`
	Avatar       string     `json:"avatar,omitempty"`
	Bio          string     `json:"bio,omitempty"`
	IsActive     bool       `json:"isActive"`
	IsVerified   bool       `json:"isVerified"`
	LastLoginAt  *time.Time `json:"lastLoginAt,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Gender       string     `json:"gender,omitempty"`
	Age          *int       `json:"age,omitempty"`
	University   string     `json:"university,omitempty"`
	Major        *string    `json:"major,omitempty"`
}

// FullName returns the user's full name
func (u *User) FullName() string {
	if u.FirstName != "" && u.LastName != "" {
		return u.FirstName + " " + u.LastName
	}
	return u.Username
}

// ToPublicUser returns a user object safe for public display
func (u *User) ToPublicUser() *PublicUser {
	return &PublicUser{
		ID:         u.ID,
		Username:   u.Username,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Avatar:     u.Avatar,
		Bio:        u.Bio,
		Gender:     u.Gender,
		Age:        u.Age,
		University: u.University,
		Major:      u.Major,
		CreatedAt:  u.CreatedAt,
	}
}

// PublicUser represents user data safe for public display
type PublicUser struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	Bio        string    `json:"bio,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Age        *int      `json:"age,omitempty"`
	University string    `json:"university,omitempty"`
	Major      *string   `json:"major,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

// CreateUserRequest represents the data needed to create a new user
type CreateUserRequest struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Password   string  `json:"password"`
	Gender     string  `json:"gender,omitempty"`
	Age        *int    `json:"age,omitempty"`
	University string  `json:"university,omitempty"`
	Major      *string `json:"major,omitempty"`
}

// UpdateUserRequest represents the data that can be updated
type UpdateUserRequest struct {
	Email      string `json:"email,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Bio        string  `json:"bio,omitempty"`
	Gender     string  `json:"gender,omitempty"`
	Age        *int    `json:"age,omitempty"`
	University string  `json:"university,omitempty"`
	Major      *string `json:"major,omitempty"`
}

// UserStats represents user statistics
type UserStats struct {
	TotalUsers    int `json:"totalUsers"`
	ActiveUsers   int `json:"activeUsers"`
	VerifiedUsers int `json:"verifiedUsers"`
}
