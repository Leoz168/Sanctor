package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"sanctor/internal/user"
)

// Service handles authentication business logic
// JWT secret (should be loaded from config/env)
var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "your-secret-key"
	}
	return secret
}

// GenerateJWT creates a JWT token for a user
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT parses and validates a JWT token
func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}
	userID, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("userId not found in token")
	}
	return userID, nil
}

type Service struct {
	repo        *Repository
	userService *user.Service
}

// NewService creates a new instance of the Service
func NewService(repo *Repository, userService *user.Service) *Service {
	return &Service{
		repo:        repo,
		userService: userService,
	}
}

// Login authenticates a user and returns a token
func (s *Service) Login(req LoginRequest) (*AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}
	// Find user by email
	u, err := s.userService.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	// Check password
	if !user.CheckPassword(req.Password, u.PasswordHash) {
		return nil, errors.New("invalid password")
	}
	// Generate JWT
	token, err := GenerateJWT(u.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	// Return response
	return &AuthResponse{
		Token:        token,
		RefreshToken: "", // TODO: implement refresh token
		ExpiresAt:    time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

// Register creates a new user and returns a token
func (s *Service) Register(req RegisterRequest) (*AuthResponse, error) {
	if req.Email == "" || req.Username == "" || req.Password == "" {
		return nil, errors.New("email, username, and password are required")
	}
	// Create user
	userReq := user.CreateUserRequest{
		Email:     req.Email,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}
	u, err := s.userService.CreateUser(userReq)
	if err != nil {
		return nil, err
	}
	// Generate JWT
	token, err := GenerateJWT(u.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	// Return response
	return &AuthResponse{
		Token:        token,
		RefreshToken: "", // TODO: implement refresh token
		ExpiresAt:    time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

// ValidateToken validates a JWT token
func (s *Service) ValidateToken(token string) (string, error) {
	return ValidateJWT(token)
}
