package middleware

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
	redisclient "sanctor/internal/redis"
	"strconv"
	"time"
)

// Logger is a middleware that logs HTTP requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log request details
		log.Printf(
			"%s %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// CORS enables CORS for all requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Authenticate is a middleware that validates JWT tokens
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate the token
		claims, err := ValidateJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user role to context
		r = r.WithContext(context.WithValue(r.Context(), "userRole", claims.Role))
		next.ServeHTTP(w, r)
	})
}

// RateLimitConfig controls Redis-backed fixed-window rate limiting.
type RateLimitConfig struct {
	Prefix      string
	Window      time.Duration
	MaxRequests int64
	FailOpen    bool
}

// DefaultRateLimitConfig returns conservative defaults for public APIs.
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		Prefix:      "rl",
		Window:      time.Minute,
		MaxRequests: 120,
		FailOpen:    true,
	}
}

// RateLimit is a middleware that limits request rate
func RateLimit(next http.Handler) http.Handler {
	return RateLimitWithRedis(nil, DefaultRateLimitConfig())(next)
}

// RateLimitWithRedis enforces per-identity, per-route limits backed by Redis.
func RateLimitWithRedis(client *redisclient.Client, cfg RateLimitConfig) func(http.Handler) http.Handler {
	if cfg.Prefix == "" {
		cfg.Prefix = "rl"
	}
	if cfg.Window <= 0 {
		cfg.Window = time.Minute
	}
	if cfg.MaxRequests <= 0 {
		cfg.MaxRequests = 120
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			if client == nil || client.Raw() == nil {
				next.ServeHTTP(w, r)
				return
			}

			identity := requestIdentity(r)
			now := time.Now().UTC()
			windowSeconds := int64(cfg.Window / time.Second)
			if windowSeconds <= 0 {
				windowSeconds = 60
			}

			bucket := now.Unix() / windowSeconds
			key := fmt.Sprintf("%s:%s:%s:%d", cfg.Prefix, r.URL.Path, identity, bucket)

			ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
			defer cancel()

			pipe := client.Raw().TxPipeline()
			countCmd := pipe.Incr(ctx, key)
			pipe.Expire(ctx, key, cfg.Window)
			if _, err := pipe.Exec(ctx); err != nil {
				if cfg.FailOpen {
					log.Printf("rate limiter redis error (fail-open): %v", err)
					next.ServeHTTP(w, r)
					return
				}
				http.Error(w, "Rate limiter unavailable", http.StatusServiceUnavailable)
				return
			}

			count := countCmd.Val()
			remaining := cfg.MaxRequests - count
			if remaining < 0 {
				remaining = 0
			}
			resetAt := ((bucket + 1) * windowSeconds)

			w.Header().Set("X-RateLimit-Limit", strconv.FormatInt(cfg.MaxRequests, 10))
			w.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(remaining, 10))
			w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(resetAt, 10))

			if count > cfg.MaxRequests {
				w.Header().Set("Retry-After", strconv.FormatInt(windowSeconds, 10))
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func requestIdentity(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if auth != "" {
		hash := sha1.Sum([]byte(auth))
		return "token:" + hex.EncodeToString(hash[:8])
	}

	ip := firstNonEmpty(
		r.Header.Get("X-Forwarded-For"),
		r.Header.Get("X-Real-IP"),
	)
	if ip != "" {
		return "ip:" + ip
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && host != "" {
		return "ip:" + host
	}

	if r.RemoteAddr != "" {
		return "ip:" + r.RemoteAddr
	}

	return "ip:unknown"
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// ValidateJWT validates a JWT token and returns the claims
func ValidateJWT(token string) (Claims, error) {
	// TODO: Implement JWT validation logic
	return Claims{}, nil
}

// Claims represents the JWT claims
type Claims struct {
	Role string `json:"role"`
}
