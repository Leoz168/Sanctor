package user

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func withTestUserService(t *testing.T, fn func(testRepo Repository, testService *Service)) {
	t.Helper()

	previousRepo := repo
	previousService := service

	testRepo := NewRepository()
	testService := NewService(testRepo)
	repo = testRepo
	service = testService

	defer func() {
		repo = previousRepo
		service = previousService
	}()

	fn(testRepo, testService)
}

func requestWithCurrentUser(r *http.Request, userID uuid.UUID) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), "userId", userID.String()))
}

func TestGetCurrentUserRequiresAuthentication(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/users/me", nil)
	recorder := httptest.NewRecorder()

	GetCurrentUser(recorder, request)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", recorder.Code)
	}
}

func TestUpdateCurrentUserRejectsInvalidEmail(t *testing.T) {
	withTestUserService(t, func(testRepo Repository, testService *Service) {
		existingUser := &User{
			ID:       uuid.New(),
			Email:    "valid@example.com",
			Username: "valid-user",
		}
		if err := testRepo.Create(existingUser); err != nil {
			t.Fatalf("failed to seed user: %v", err)
		}

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/users/me",
			bytes.NewBufferString(`{"email":"bad-email"}`),
		)
		request.Header.Set("Content-Type", "application/json")
		request = requestWithCurrentUser(request, existingUser.ID)
		recorder := httptest.NewRecorder()

		UpdateCurrentUser(recorder, request)

		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", recorder.Code)
		}
	})
}

func TestDeleteCurrentUserDeletesAuthenticatedUser(t *testing.T) {
	withTestUserService(t, func(testRepo Repository, testService *Service) {
		existingUser := &User{
			ID:       uuid.New(),
			Email:    "delete@example.com",
			Username: "delete-user",
		}
		if err := testRepo.Create(existingUser); err != nil {
			t.Fatalf("failed to seed user: %v", err)
		}

		request := httptest.NewRequest(http.MethodDelete, "/api/users/me", nil)
		request = requestWithCurrentUser(request, existingUser.ID)
		recorder := httptest.NewRecorder()

		DeleteCurrentUser(recorder, request)

		if recorder.Code != http.StatusNoContent {
			t.Fatalf("expected 204, got %d", recorder.Code)
		}

		if _, err := testService.GetUser(existingUser.ID); err == nil {
			t.Fatalf("expected user to be deleted")
		}
	})
}
