package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// VerifyPassword checks if the provided password matches the user's password
func (s *Service) VerifyPassword(userID uuid.UUID, password string) (bool, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return false, err
	}

	return CheckPassword(password, user.PasswordHash), nil
}

// ChangePassword updates a user's password
func (s *Service) ChangePassword(userID uuid.UUID, oldPassword, newPassword string) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Verify old password
	if !CheckPassword(oldPassword, user.PasswordHash) {
		return errors.New("invalid current password")
	}

	// Validate new password
	if len(newPassword) < 8 {
		return errors.New("new password must be at least 8 characters")
	}

	// Hash new password
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.PasswordHash = hashedPassword
	return s.repo.Update(user)
}

// FindByEmail retrieves a user by email
func (s *Service) FindByEmail(email string) (*User, error) {
	return s.repo.FindByEmail(email)
}

// FindByUsername retrieves a user by username
func (s *Service) FindByUsername(username string) (*User, error) {
	return s.repo.FindByUsername(username)
}

// FindByGoogleSub retrieves a user by Google subject identifier.
func (s *Service) FindByGoogleSub(sub string) (*User, error) {
	return s.repo.FindByGoogleSub(sub)
}

// ExistsByUsername checks if a username already exists.
func (s *Service) ExistsByUsername(username string) bool {
	return s.repo.ExistsByUsername(username)
}

// LinkGoogleSub attaches a Google subject to an existing user.
func (s *Service) LinkGoogleSub(userID uuid.UUID, sub string, emailVerified bool, picture string, firstName string, lastName string) (*User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if user.GoogleSub != nil && *user.GoogleSub != sub {
		return nil, errors.New("google account already linked")
	}
	user.GoogleSub = &sub
	if emailVerified {
		user.IsVerified = true
	}
	if user.Avatar == "" && picture != "" {
		user.Avatar = picture
	}
	if user.FirstName == "" && firstName != "" {
		user.FirstName = firstName
	}
	if user.LastName == "" && lastName != "" {
		user.LastName = lastName
	}
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
