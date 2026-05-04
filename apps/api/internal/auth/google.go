package auth

import (
	"context"
	"errors"
	"os"
	"strings"

	"google.golang.org/api/idtoken"
)

var (
	ErrGoogleTokenRequired   = errors.New("google id token is required")
	ErrGoogleTokenInvalid    = errors.New("google id token is invalid")
	ErrGoogleEmailMissing    = errors.New("google account email is missing")
	ErrGoogleEmailUnverified = errors.New("google account email is not verified")
	ErrGoogleSubMissing      = errors.New("google account subject is missing")
	ErrGoogleEmailMismatch   = errors.New("google account email does not match")
	ErrGoogleAccountInUse    = errors.New("google account is already linked")
)

type googleTokenValidator interface {
	Validate(ctx context.Context, token, audience string) (*GoogleClaims, error)
}

type googleIDTokenValidator struct{}

func (v googleIDTokenValidator) Validate(ctx context.Context, token, audience string) (*GoogleClaims, error) {
	if strings.TrimSpace(audience) == "" {
		return nil, errors.New("google client id is required")
	}
	payload, err := idtoken.Validate(ctx, token, audience)
	if err != nil {
		return nil, err
	}
	return googleClaimsFromPayload(payload), nil
}

func googleClientID() string {
	return strings.TrimSpace(os.Getenv("GOOGLE_CLIENT_ID"))
}

type GoogleClaims struct {
	Subject       string
	Email         string
	EmailVerified bool
	FirstName     string
	LastName      string
	FullName      string
	Picture       string
}

func googleClaimsFromPayload(payload *idtoken.Payload) *GoogleClaims {
	claims := &GoogleClaims{}
	if payload == nil {
		return claims
	}
	claims.Subject = payload.Subject

	if value, ok := payload.Claims["email"].(string); ok {
		claims.Email = value
	}
	if value, ok := payload.Claims["email_verified"].(bool); ok {
		claims.EmailVerified = value
	}
	if value, ok := payload.Claims["given_name"].(string); ok {
		claims.FirstName = value
	}
	if value, ok := payload.Claims["family_name"].(string); ok {
		claims.LastName = value
	}
	if value, ok := payload.Claims["name"].(string); ok {
		claims.FullName = value
	}
	if value, ok := payload.Claims["picture"].(string); ok {
		claims.Picture = value
	}

	return claims
}
