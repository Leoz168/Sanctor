package user

import (
    "errors"
    "testing"
    "github.com/google/uuid"
)

type userStub struct{
    created *User
    findErr error
    existsEmail bool
    existsUsername bool
    blacklisted bool
}
func (r *userStub) Create(u *User) error { r.created = u; return nil }
func (r *userStub) FindByID(id uuid.UUID) (*User, error) { if r.findErr!=nil { return nil, r.findErr }; return r.created, nil }
func (r *userStub) FindAll() []*User { return []*User{} }
func (r *userStub) Update(u *User) error { r.created = u; return nil }
func (r *userStub) Delete(id uuid.UUID) error { return nil }
func (r *userStub) ExistsByEmail(email string) bool { return r.existsEmail }
func (r *userStub) IsEmailBlacklisted(email string) (bool, error) { return r.blacklisted, nil }
func (r *userStub) ExistsByUsername(username string) bool { return r.existsUsername }
func (r *userStub) FindByEmail(email string) (*User, error) { return r.created, nil }
func (r *userStub) FindByUsername(username string) (*User, error) { return r.created, nil }
func (r *userStub) FindByGoogleSub(sub string) (*User, error) { return r.created, nil }

func TestCreateUserValidationAndDuplicates(t *testing.T){
    svc := NewService(&userStub{})
    _, err := svc.CreateUser(CreateUserRequest{Email: "", Username: "u", Password: "12345678"})
    if err==nil { t.Fatalf("expected email required") }

    // invalid email
    _, err = svc.CreateUser(CreateUserRequest{Email: "noat", Username: "usr", Password: "12345678"})
    if err==nil { t.Fatalf("expected invalid email") }

    // blacklisted
    repo := &userStub{blacklisted:true}
    svc = NewService(repo)
    _, err = svc.CreateUser(CreateUserRequest{Email: "a@b.com", Username: "usr", Password: "12345678"})
    if err==nil || !errors.Is(err, ErrEmailBlacklisted) { t.Fatalf("expected email blacklisted") }
}

func TestGetUserValidation(t *testing.T){
    svc := NewService(&userStub{})
    _, err := svc.GetUser(uuid.Nil)
    if err==nil || err.Error()!="user ID is required" { t.Fatalf("expected id required") }
}