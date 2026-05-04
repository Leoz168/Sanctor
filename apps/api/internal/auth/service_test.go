package auth

import (
    "context"
    "errors"
    "testing"
    "time"

    "github.com/google/uuid"
    "sanctor/internal/user"
)

type stubGoogleValidator struct {
    claims *GoogleClaims
    err    error
}

func (s stubGoogleValidator) Validate(ctx context.Context, token, audience string) (*GoogleClaims, error) {
    if s.err != nil {
        return nil, s.err
    }
    if audience == "" {
        return nil, errors.New("missing audience")
    }
    return s.claims, nil
}

func TestGenerateAndValidateJWT(t *testing.T) {
    token, err := GenerateJWT("test-user-id")
    if err != nil {
        t.Fatalf("generate jwt error: %v", err)
    }
    uid, err := ValidateJWT(token)
    if err != nil {
        t.Fatalf("validate jwt error: %v", err)
    }
    if uid != "test-user-id" {
        t.Fatalf("expected user id match, got %s", uid)
    }
}

func TestLoginFlows(t *testing.T) {
    repo := user.NewRepository()
    us := user.NewService(repo)

    // create a user
    hashed, _ := user.HashPassword("strongpassword")
    u := &user.User{ID: uuid.New(), Email: "a@b.com", Username: "u1", PasswordHash: hashed}
    if err := repo.Create(u); err != nil {
        t.Fatalf("create user: %v", err)
    }

    svc := NewService(nil, us)

    // missing fields
    if _, err := svc.Login(LoginRequest{Email: "", Password: ""}); err == nil {
        t.Fatalf("expected validation error")
    }

    // wrong password
    if _, err := svc.Login(LoginRequest{Email: "a@b.com", Password: "bad"}); err == nil {
        t.Fatalf("expected invalid password")
    }

    // blacklisted
    u.IsBlacklisted = true
    if err := repo.Update(u); err != nil { t.Fatalf("update user: %v", err) }
    if _, err := svc.Login(LoginRequest{Email: "a@b.com", Password: "strongpassword"}); err == nil {
        t.Fatalf("expected blacklisted error")
    }

    // un-blacklist and login success
    u.IsBlacklisted = false
    if err := repo.Update(u); err != nil { t.Fatalf("update user: %v", err) }
    resp, err := svc.Login(LoginRequest{Email: "a@b.com", Password: "strongpassword"})
    if err != nil { t.Fatalf("expected login success, got %v", err) }
    if resp.Token == "" { t.Fatalf("expected token") }
    // token should validate
    if _, err := ValidateJWT(resp.Token); err != nil { t.Fatalf("token validate: %v", err) }
    // expires at reasonable
    if _, err := time.Parse(time.RFC3339, resp.ExpiresAt); err != nil { t.Fatalf("invalid expires format: %v", err) }
}

func TestLoginWithGoogleCreatesUser(t *testing.T) {
    repo := user.NewRepository()
    us := user.NewService(repo)
    validator := stubGoogleValidator{
        claims: &GoogleClaims{
            Subject:       "google-sub",
            Email:         "test@example.com",
            EmailVerified: true,
            FirstName:     "Test",
            LastName:      "User",
            Picture:       "http://example.com/pic.png",
        },
    }
    service := NewServiceWithGoogleValidator(nil, us, validator, "client-id")

    resp, err := service.LoginWithGoogle(context.Background(), GoogleAuthRequest{IDToken: "token"})
    if err != nil {
        t.Fatalf("expected success, got %v", err)
    }
    if resp.Token == "" {
        t.Fatalf("expected token")
    }

    created, err := us.FindByEmail("test@example.com")
    if err != nil {
        t.Fatalf("expected user created, got %v", err)
    }
    if created.GoogleSub == nil || *created.GoogleSub != "google-sub" {
        t.Fatalf("expected google sub to be set")
    }
    if !created.IsVerified {
        t.Fatalf("expected user to be verified")
    }
}

func TestLoginWithGoogleLinksExistingByEmail(t *testing.T) {
    repo := user.NewRepository()
    us := user.NewService(repo)
    created, err := us.CreateUser(user.CreateUserRequest{
        Email:    "link@example.com",
        Username: "linkuser",
        Password: "password123",
    })
    if err != nil {
        t.Fatalf("create user: %v", err)
    }

    validator := stubGoogleValidator{
        claims: &GoogleClaims{Subject: "google-link", Email: "link@example.com", EmailVerified: true},
    }
    service := NewServiceWithGoogleValidator(nil, us, validator, "client-id")

    _, err = service.LoginWithGoogle(context.Background(), GoogleAuthRequest{IDToken: "token"})
    if err != nil {
        t.Fatalf("expected login to succeed, got %v", err)
    }

    linked, err := us.FindByGoogleSub("google-link")
    if err != nil {
        t.Fatalf("expected linked user, got %v", err)
    }
    if linked.ID != created.ID {
        t.Fatalf("expected same user linked")
    }
}

func TestLinkGoogleAccount(t *testing.T) {
    repo := user.NewRepository()
    us := user.NewService(repo)
    created, err := us.CreateUser(user.CreateUserRequest{
        Email:    "link2@example.com",
        Username: "linkuser2",
        Password: "password123",
    })
    if err != nil {
        t.Fatalf("create user: %v", err)
    }

    validator := stubGoogleValidator{
        claims: &GoogleClaims{Subject: "google-link2", Email: "link2@example.com", EmailVerified: true},
    }
    service := NewServiceWithGoogleValidator(nil, us, validator, "client-id")

    if err := service.LinkGoogleAccount(context.Background(), created.ID, GoogleAuthRequest{IDToken: "token"}); err != nil {
        t.Fatalf("expected link success, got %v", err)
    }

    linked, err := us.FindByGoogleSub("google-link2")
    if err != nil {
        t.Fatalf("expected linked user, got %v", err)
    }
    if linked.ID != created.ID {
        t.Fatalf("expected same user linked")
    }
}

func TestLoginWithGoogleRejectsUnverifiedEmail(t *testing.T) {
    repo := user.NewRepository()
    us := user.NewService(repo)
    validator := stubGoogleValidator{
        claims: &GoogleClaims{Subject: "sub", Email: "nope@example.com", EmailVerified: false},
    }
    service := NewServiceWithGoogleValidator(nil, us, validator, "client-id")

    _, err := service.LoginWithGoogle(context.Background(), GoogleAuthRequest{IDToken: "token"})
    if err == nil || !errors.Is(err, ErrGoogleEmailUnverified) {
        t.Fatalf("expected unverified error")
    }
}
