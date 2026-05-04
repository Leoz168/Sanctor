package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticateMissingToken(t *testing.T) {
	h := Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rr.Code)
	}
}

func TestCORSHeaders(t *testing.T) {
	called := false
	h := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(rr, req)

	if !called {
		t.Fatalf("expected handler to be called")
	}
	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Fatalf("missing CORS header")
	}
}
