package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"sanctor/internal/auth"
)

func TestRequestIdentityAuthorization(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer token")

	sum := sha1.Sum([]byte("Bearer token"))
	want := "token:" + hex.EncodeToString(sum[:8])
	if got := requestIdentity(req); got != want {
		t.Fatalf("expected %s, got %s", want, got)
	}
}

func TestRequestIdentityForwardedFor(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Forwarded-For", "1.2.3.4, 5.6.7.8")
	if got := requestIdentity(req); got != "ip:1.2.3.4" {
		t.Fatalf("expected ip from xff, got %s", got)
	}
}

func TestResolveRateLimitConfigLongestPrefix(t *testing.T) {
	defaultCfg := RateLimitConfig{Prefix: "rl", Window: time.Second, MaxRequests: 10}
	policies := map[string]RateLimitConfig{
		"/api":    {Prefix: "api", Window: time.Minute, MaxRequests: 20},
		"/api/v1": {Prefix: "v1", Window: time.Minute, MaxRequests: 5},
	}

	cfg := resolveRateLimitConfig("/api/v1/users", policies, defaultCfg)
	if cfg.Prefix != "v1" {
		t.Fatalf("expected v1 policy, got %s", cfg.Prefix)
	}
}

func TestRateLimitFailOpen(t *testing.T) {
	called := false
	h := RateLimitWithRedis(nil, RateLimitConfig{FailOpen: true, Window: time.Second, MaxRequests: 1})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(rr, req)

	if !called {
		t.Fatalf("expected handler to be called")
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestRateLimitFailClosed(t *testing.T) {
	called := false
	h := RateLimitWithRedis(nil, RateLimitConfig{FailOpen: false, Window: time.Second, MaxRequests: 1})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(rr, req)

	if called {
		t.Fatalf("expected handler not to be called")
	}
	if rr.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rr.Code)
	}
}

func TestCORSOptionsShortCircuit(t *testing.T) {
	called := false
	h := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	h.ServeHTTP(rr, req)

	if called {
		t.Fatalf("expected handler not to be called")
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestAuthenticateValidToken(t *testing.T) {
	var got string
	token, err := auth.GenerateJWT("user-1")
	if err != nil {
		t.Fatalf("generate token: %v", err)
	}

	h := Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := r.Context().Value("userId")
		if value != nil {
			got = value.(string)
		}
		w.WriteHeader(http.StatusOK)
	}))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if got != "user-1" {
		t.Fatalf("expected user id, got %s", got)
	}
}
