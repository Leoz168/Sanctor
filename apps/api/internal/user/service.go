package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Service handles business logic for user operations
type Service struct {
	repo Repository
}

// NewService creates a new user service
func NewService(repo Repository) *Service {
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

	blacklisted, err := s.repo.IsEmailBlacklisted(req.Email)
	if err != nil {
		return nil, err
	}
	if blacklisted {
		return nil, ErrEmailBlacklisted
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

	var institutionID *uuid.UUID
	if req.InstitutionID != nil && *req.InstitutionID != "" {
		parsed, err := uuid.Parse(*req.InstitutionID)
		if err != nil {
			return nil, errors.New("invalid institution ID format")
		}
		institutionID = &parsed
	}

	// Create user
	user := &User{
		ID:            uuid.New(),
		Email:         req.Email,
		Username:      req.Username,
		PasswordHash:  hashedPassword,
		Gender:        req.Gender,
		Age:           req.Age,
		InstitutionID: institutionID,
		Major:         req.Major,
		IsActive:      true,
		IsVerified:    false,
		IsBlacklisted: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(id uuid.UUID) (*User, error) {
	if id == uuid.Nil {
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
func (s *Service) UpdateUser(id uuid.UUID, req UpdateUserRequest) (*User, error) {
	if id == uuid.Nil {
		return nil, errors.New("user ID is required")
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update fields if provided
	if req.Username != "" && req.Username != user.Username {
		if err := ValidateUsername(req.Username); err != nil {
			return nil, err
		}
		if s.repo.ExistsByUsername(req.Username) {
			return nil, errors.New("username already taken")
		}
		user.Username = req.Username
	}
	if req.Email != "" {
		if !ValidateEmail(req.Email) {
			return nil, errors.New("invalid email format")
		}
		if req.Email != user.Email {
			blacklisted, err := s.repo.IsEmailBlacklisted(req.Email)
			if err != nil {
				return nil, err
			}
			if blacklisted {
				return nil, ErrEmailBlacklisted
			}
			if s.repo.ExistsByEmail(req.Email) {
				return nil, errors.New("user with this email already exists")
			}
		}
		user.Email = req.Email
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
	if req.InstitutionID != nil {
		if *req.InstitutionID == "" {
			user.InstitutionID = nil
		} else {
			parsed, err := uuid.Parse(*req.InstitutionID)
			if err != nil {
				return nil, errors.New("invalid institution ID format")
			}
			user.InstitutionID = &parsed
		}
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

// GetCurrentUser retrieves the authenticated user.
func (s *Service) GetCurrentUser(id uuid.UUID) (*User, error) {
	return s.GetUser(id)
}

// UpdateCurrentUser updates the authenticated user.
func (s *Service) UpdateCurrentUser(id uuid.UUID, req UpdateUserRequest) (*User, error) {
	return s.UpdateUser(id, req)
}

// DeleteCurrentUser deletes the authenticated user.
func (s *Service) DeleteCurrentUser(id uuid.UUID) error {
	return s.DeleteUser(id)
}

// DeleteUser deletes a user by ID
func (s *Service) DeleteUser(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("user ID is required")
	}

	if err := s.repo.Delete(id); err != nil {
		return errors.New("user not found")
	}

	return nil
}
