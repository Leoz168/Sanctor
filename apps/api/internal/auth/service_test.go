package auth

import (
    "testing"
    "time"

    "github.com/google/uuid"
    "sanctor/internal/user"
)

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
