package user

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

type userStub struct {
	created        *User
	findErr        error
	existsEmail    bool
	existsUsername bool
	blacklisted    bool
}

func (r *userStub) Create(u *User) error { r.created = u; return nil }
func (r *userStub) FindByID(id uuid.UUID) (*User, error) {
	if r.findErr != nil {
		return nil, r.findErr
	}
	return r.created, nil
}
func (r *userStub) FindAll() []*User                              { return []*User{} }
func (r *userStub) Update(u *User) error                          { r.created = u; return nil }
func (r *userStub) Delete(id uuid.UUID) error                     { return nil }
func (r *userStub) ExistsByEmail(email string) bool               { return r.existsEmail }
func (r *userStub) IsEmailBlacklisted(email string) (bool, error) { return r.blacklisted, nil }
func (r *userStub) ExistsByUsername(username string) bool         { return r.existsUsername }
func (r *userStub) FindByEmail(email string) (*User, error)       { return r.created, nil }
func (r *userStub) FindByUsername(username string) (*User, error) { return r.created, nil }
func (r *userStub) FindByGoogleSub(sub string) (*User, error)     { return r.created, nil }

func TestCreateUserValidationAndDuplicates(t *testing.T) {
	svc := NewService(&userStub{})
	_, err := svc.CreateUser(CreateUserRequest{Email: "", Username: "u", Password: "12345678"})
	if err == nil {
		t.Fatalf("expected email required")
	}

	_, err = svc.CreateUser(CreateUserRequest{Email: "noat", Username: "usr", Password: "12345678"})
	if err == nil {
		t.Fatalf("expected invalid email")
	}

	repo := &userStub{blacklisted: true}
	svc = NewService(repo)
	_, err = svc.CreateUser(CreateUserRequest{Email: "a@b.com", Username: "usr", Password: "12345678"})
	if err == nil || !errors.Is(err, ErrEmailBlacklisted) {
		t.Fatalf("expected email blacklisted")
	}
}

func TestGetUserValidation(t *testing.T) {
	svc := NewService(&userStub{})
	_, err := svc.GetUser(uuid.Nil)
	if err == nil || err.Error() != "user ID is required" {
		t.Fatalf("expected id required")
	}
}

func TestUpdateUserValidatesEmailFormat(t *testing.T) {
	existing := &User{
		ID:       uuid.New(),
		Email:    "before@example.com",
		Username: "before",
	}
	svc := NewService(&userStub{created: existing})

	_, err := svc.UpdateUser(existing.ID, UpdateUserRequest{Email: "bad-email"})
	if err == nil || err.Error() != "invalid email format" {
		t.Fatalf("expected invalid email format, got %v", err)
	}
}

func TestUpdateUserValidatesUsername(t *testing.T) {
	existing := &User{
		ID:       uuid.New(),
		Email:    "before@example.com",
		Username: "before",
	}
	svc := NewService(&userStub{created: existing, existsUsername: true})

	_, err := svc.UpdateUser(existing.ID, UpdateUserRequest{Username: "taken"})
	if err == nil || err.Error() != "username already taken" {
		t.Fatalf("expected duplicate username error, got %v", err)
	}
}

func TestCurrentUserMethodsDelegateToGenericCRUD(t *testing.T) {
	existing := &User{
		ID:       uuid.New(),
		Email:    "before@example.com",
		Username: "before",
	}
	repo := &userStub{created: existing}
	svc := NewService(repo)

	currentUser, err := svc.GetCurrentUser(existing.ID)
	if err != nil || currentUser.ID != existing.ID {
		t.Fatalf("expected current user lookup to succeed, got user=%v err=%v", currentUser, err)
	}

	updatedUser, err := svc.UpdateCurrentUser(existing.ID, UpdateUserRequest{Bio: "updated bio"})
	if err != nil || updatedUser.Bio != "updated bio" {
		t.Fatalf("expected current user update to succeed, got user=%v err=%v", updatedUser, err)
	}

	if err := svc.DeleteCurrentUser(existing.ID); err != nil {
		t.Fatalf("expected current user delete to succeed, got %v", err)
	}
}

func TestSearchUsersExcludesCurrentUserAndMatchesByUsername(t *testing.T) {
	repo := &userStub{}
	svc := NewService(repo)

	currentUser := &User{ID: uuid.New(), Email: "self@example.com", Username: "self"}
	otherUser := &User{ID: uuid.New(), Email: "alex@example.com", Username: "alexrivera"}

	repo.created = currentUser
	memRepo := NewRepository().(*InMemoryRepository)
	_ = memRepo.Create(currentUser)
	_ = memRepo.Create(otherUser)
	svc = NewService(memRepo)

	results, err := svc.SearchUsers("alex", currentUser.ID)
	if err != nil {
		t.Fatalf("expected search to succeed, got %v", err)
	}
	if len(results) != 1 || results[0].ID != otherUser.ID {
		t.Fatalf("expected only alex to be returned, got %+v", results)
	}
}
