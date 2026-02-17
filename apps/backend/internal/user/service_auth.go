package user

import "errors"

// VerifyPassword checks if the provided password matches the user's password
func (s *Service) VerifyPassword(userID, password string) (bool, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return false, err
	}
	
	return CheckPassword(password, user.PasswordHash), nil
}

// ChangePassword updates a user's password
func (s *Service) ChangePassword(userID, oldPassword, newPassword string) error {
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
