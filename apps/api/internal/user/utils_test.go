package user

import (
	"errors"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		expectError   bool
		expectedError error
	}{
		{
			name:          "Password too short",
			password:      "short",
			expectError:   true,
			expectedError: errors.New("password must be at least 8 characters"),
		},
		{
			name:          "Password meets length requirement",
			password:      "longenough",
			expectError:   false,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.password)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}