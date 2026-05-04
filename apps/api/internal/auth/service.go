package auth

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	googleValidator googleTokenValidator
	googleAudience  string
}

// NewService creates a new instance of the Service
func NewService(repo *Repository, userService *user.Service) *Service {
	return NewServiceWithGoogleValidator(repo, userService, googleIDTokenValidator{}, googleClientID())
}

// NewServiceWithGoogleValidator creates a new service with a custom Google validator (for tests).
func NewServiceWithGoogleValidator(repo *Repository, userService *user.Service, validator googleTokenValidator, audience string) *Service {
	if validator == nil {
		validator = googleIDTokenValidator{}
	}
	return &Service{
		repo:            repo,
		userService:     userService,
		googleValidator: validator,
		googleAudience:  strings.TrimSpace(audience),
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
	if u.IsBlacklisted {
		return nil, user.ErrUserBlacklisted
	}
	// Check password
	if !user.CheckPassword(req.Password, u.PasswordHash) {
		return nil, errors.New("invalid password")
	}
	return issueAuthResponse(u.ID.String())
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
	return issueAuthResponse(u.ID.String())
}

// ValidateToken validates a JWT token
func (s *Service) ValidateToken(token string) (string, error) {
	return ValidateJWT(token)
}

// LoginWithGoogle registers or logs in a user using a Google ID token.
func (s *Service) LoginWithGoogle(ctx context.Context, req GoogleAuthRequest) (*AuthResponse, error) {
	token := strings.TrimSpace(req.IDToken)
	if token == "" {
		return nil, ErrGoogleTokenRequired
	}

	claims, err := s.googleValidator.Validate(ctx, token, s.googleAudience)
	if err != nil {
		return nil, ErrGoogleTokenInvalid
	}
	if claims.Subject == "" {
		return nil, ErrGoogleSubMissing
	}
	if claims.Email == "" {
		return nil, ErrGoogleEmailMissing
	}
	if !claims.EmailVerified {
		return nil, ErrGoogleEmailUnverified
	}

	if existing, err := s.userService.FindByGoogleSub(claims.Subject); err == nil && existing != nil {
		if existing.IsBlacklisted {
			return nil, user.ErrUserBlacklisted
		}
		return issueAuthResponse(existing.ID.String())
	}

	if existingByEmail, err := s.userService.FindByEmail(claims.Email); err == nil && existingByEmail != nil {
		if existingByEmail.IsBlacklisted {
			return nil, user.ErrUserBlacklisted
		}
		if existingByEmail.GoogleSub != nil && *existingByEmail.GoogleSub != claims.Subject {
			return nil, ErrGoogleAccountInUse
		}
		if _, err := s.userService.LinkGoogleSub(existingByEmail.ID, claims.Subject, claims.EmailVerified, claims.Picture, claims.FirstName, claims.LastName); err != nil {
			if err.Error() == "google account already linked" {
				return nil, ErrGoogleAccountInUse
			}
			return nil, err
		}
		return issueAuthResponse(existingByEmail.ID.String())
	}

	username := generateUsername(claims.Email, s.userService)
	password := uuid.NewString()
	userReq := user.CreateUserRequest{
		Email:     claims.Email,
		Username:  username,
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		Password:  password,
	}

	createdUser, err := s.userService.CreateUser(userReq)
	if err != nil {
		return nil, err
	}
	if _, err := s.userService.LinkGoogleSub(createdUser.ID, claims.Subject, claims.EmailVerified, claims.Picture, claims.FirstName, claims.LastName); err != nil {
		if err.Error() == "google account already linked" {
			return nil, ErrGoogleAccountInUse
		}
		return nil, err
	}

	return issueAuthResponse(createdUser.ID.String())
}

// LinkGoogleAccount links a Google identity to an existing user.
func (s *Service) LinkGoogleAccount(ctx context.Context, userID uuid.UUID, req GoogleAuthRequest) error {
	token := strings.TrimSpace(req.IDToken)
	if token == "" {
		return ErrGoogleTokenRequired
	}

	claims, err := s.googleValidator.Validate(ctx, token, s.googleAudience)
	if err != nil {
		return ErrGoogleTokenInvalid
	}
	if claims.Subject == "" {
		return ErrGoogleSubMissing
	}
	if claims.Email == "" {
		return ErrGoogleEmailMissing
	}
	if !claims.EmailVerified {
		return ErrGoogleEmailUnverified
	}

	if existing, err := s.userService.FindByGoogleSub(claims.Subject); err == nil && existing != nil {
		if existing.ID != userID {
			return ErrGoogleAccountInUse
		}
		return nil
	}

	currentUser, err := s.userService.GetUser(userID)
	if err != nil {
		return err
	}
	if !strings.EqualFold(currentUser.Email, claims.Email) {
		return ErrGoogleEmailMismatch
	}

	_, err = s.userService.LinkGoogleSub(userID, claims.Subject, claims.EmailVerified, claims.Picture, claims.FirstName, claims.LastName)
	if err != nil && err.Error() == "google account already linked" {
		return ErrGoogleAccountInUse
	}
	return err
}

func issueAuthResponse(userID string) (*AuthResponse, error) {
	token, err := GenerateJWT(userID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &AuthResponse{
		Token:        token,
		RefreshToken: "",
		ExpiresAt:    time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

func generateUsername(email string, users *user.Service) string {
	base := strings.TrimSpace(strings.Split(email, "@")[0])
	if len(base) < 3 {
		base = "user"
	}
	base = trimUsername(base)
	if !users.ExistsByUsername(base) {
		return base
	}
	for i := 0; i < 5; i++ {
		suffix := uuid.NewString()[:4]
		candidate := trimUsername(base + suffix)
		if !users.ExistsByUsername(candidate) {
			return candidate
		}
	}
	return trimUsername(base + uuid.NewString()[:6])
}

func trimUsername(value string) string {
	value = strings.TrimSpace(value)
	if len(value) > 20 {
		return value[:20]
	}
	return value
}
