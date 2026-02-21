package user

import (
	"Sanctor/apps/api/internal/database"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Service handles business logic for user operations
type Service struct {
	repo *Repository
}

// NewService creates a new user service
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateUser creates a new user with validation
func (s *Service) CreateUser(req CreateUserRequest) (*User, error) {
	// Validate input
	if req.Email == "" || req.Username == "" {
		return nil, errors.New("email and username are required")
	}

	if !ValidateEmail(req.Email) {
		return nil, errors.New("invalid email format")
	}

	if err := ValidateUsername(req.Username); err != nil {
		return nil, err
	}

	if req.Password == "" || len(req.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	// Check if user already exists
	if s.repo.ExistsByEmail(req.Email) {
		return nil, errors.New("user with this email already exists")
	}

	if s.repo.ExistsByUsername(req.Username) {
		return nil, errors.New("username already taken")
	}

	// Hash password
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := &User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		Username:     req.Username,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PasswordHash: hashedPassword,
		Gender:       req.Gender,
		Age:          req.Age,
		University:   req.University,
		Major:        req.Major,
		IsActive:     true,
		IsVerified:   false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(id string) (*User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetAllUsers retrieves all users
func (s *Service) GetAllUsers() ([]*User, error) {
	return s.repo.FindAll(), nil
}

// UpdateUser updates an existing user
func (s *Service) UpdateUser(id string, req UpdateUserRequest) (*User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update fields if provided
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	if req.Age != nil {
		user.Age = req.Age
	}
	if req.University != "" {
		user.University = req.University
	}
	if req.Major != nil {
		user.Major = req.Major
	}
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (s *Service) DeleteUser(id string) error {
	if id == "" {
		return errors.New("user ID is required")
	}

	if err := s.repo.Delete(id); err != nil {
		return errors.New("user not found")
	}

	return nil
}

func CreateUser(user *User) error {
	return database.DB.Create(user).Error
}

func GetUserByID(id uint) (*User, error) {
	var user User
	err := database.DB.First(&user, id).Error
	return &user, err
}
