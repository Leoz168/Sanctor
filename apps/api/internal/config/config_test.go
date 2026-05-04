package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	keys := []string{
		"PORT",
		"HOST",
		"GO_ENV",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"REDIS_URL",
		"REDIS_HOST",
		"REDIS_PORT",
		"REDIS_USERNAME",
		"REDIS_PASSWORD",
		"REDIS_DB",
		"REDIS_SENTINEL_ENABLED",
		"REDIS_MASTER_NAME",
		"REDIS_SENTINEL_ADDRS",
		"REDIS_SENTINEL_USERNAME",
		"REDIS_SENTINEL_PASSWORD",
		"JWT_SECRET",
		"TOKEN_EXPIRY",
		"REFRESH_EXPIRY",
	}
	for _, key := range keys {
		t.Setenv(key, "")
	}

	// Set invalid values to force defaults on int/bool parsing.
	t.Setenv("DB_PORT", "not-int")
	t.Setenv("REDIS_PORT", "nope")
	t.Setenv("REDIS_DB", "nan")
	t.Setenv("TOKEN_EXPIRY", "bad")
	t.Setenv("REFRESH_EXPIRY", "bad")
	t.Setenv("REDIS_SENTINEL_ENABLED", "not-bool")

	cfg := Load()
	if cfg.Server.Port != "8080" {
		t.Fatalf("expected default port 8080, got %s", cfg.Server.Port)
	}
	if cfg.Server.Host != "0.0.0.0" {
		t.Fatalf("expected default host 0.0.0.0, got %s", cfg.Server.Host)
	}
	if cfg.Server.Env != "development" {
		t.Fatalf("expected default env development, got %s", cfg.Server.Env)
	}
	if cfg.Database.Port != 5432 {
		t.Fatalf("expected default db port 5432, got %d", cfg.Database.Port)
	}
	if cfg.Redis.DB != 0 {
		t.Fatalf("expected default redis db 0, got %d", cfg.Redis.DB)
	}
	if cfg.Redis.SentinelEnabled {
		t.Fatalf("expected default sentinel disabled")
	}
	if cfg.Auth.TokenExpiry != 24 {
		t.Fatalf("expected default token expiry 24, got %d", cfg.Auth.TokenExpiry)
	}
	if cfg.Auth.RefreshExpiry != 7 {
		t.Fatalf("expected default refresh expiry 7, got %d", cfg.Auth.RefreshExpiry)
	}
}

func TestLoadOverrides(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("HOST", "127.0.0.1")
	t.Setenv("GO_ENV", "production")
	t.Setenv("DB_HOST", "db")
	t.Setenv("DB_PORT", "5433")
	t.Setenv("DB_USER", "admin")
	t.Setenv("DB_PASSWORD", "secret")
	t.Setenv("DB_NAME", "sanctor_test")
	t.Setenv("REDIS_URL", "redis://redis:6379/0")
	t.Setenv("REDIS_PORT", "6380")
	t.Setenv("REDIS_DB", "2")
	t.Setenv("REDIS_SENTINEL_ENABLED", "true")
	t.Setenv("JWT_SECRET", "custom")
	t.Setenv("TOKEN_EXPIRY", "12")
	t.Setenv("REFRESH_EXPIRY", "3")

	cfg := Load()
	if cfg.Server.Port != "9090" {
		t.Fatalf("expected port override 9090, got %s", cfg.Server.Port)
	}
	if cfg.Database.Port != 5433 {
		t.Fatalf("expected db port override 5433, got %d", cfg.Database.Port)
	}
	if cfg.Redis.DB != 2 {
		t.Fatalf("expected redis db override 2, got %d", cfg.Redis.DB)
	}
	if !cfg.Redis.SentinelEnabled {
		t.Fatalf("expected sentinel enabled")
	}
	if cfg.Auth.JWTSecret != "custom" {
		t.Fatalf("expected jwt secret override")
	}
	if cfg.Auth.TokenExpiry != 12 || cfg.Auth.RefreshExpiry != 3 {
		t.Fatalf("expected expiry overrides")
	}
}
