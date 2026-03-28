package user

import (
	"errors"
	"testing"
)

func TestValidateNewPassword(t *testing.T) {
	tests := []struct {
		name          string
		newPassword   string
		expectError   bool
		expectedError error
	}{
		{
			name:          "New password too short",
			newPassword:   "short",
			expectError:   true,
			expectedError: errors.New("new password must be at least 8 characters"),
		},
		{
			name:          "New password meets length requirement",
			newPassword:   "longenough",
			expectError:   false,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNewPassword(tt.newPassword)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}