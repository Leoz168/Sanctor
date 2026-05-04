package database

import (
	"net/url"
	"testing"
)

func TestEnsureSSLModeAddsRequire(t *testing.T) {
	input := "postgres://user:pass@localhost:5432/db"
	out := ensureSSLMode(input)
	parsed, err := url.Parse(out)
	if err != nil {
		t.Fatalf("parse output: %v", err)
	}
	if got := parsed.Query().Get("sslmode"); got != "require" {
		t.Fatalf("expected sslmode=require, got %s", got)
	}
}

func TestEnsureSSLModeKeepsExisting(t *testing.T) {
	input := "postgres://user:pass@localhost:5432/db?sslmode=disable"
	out := ensureSSLMode(input)
	parsed, err := url.Parse(out)
	if err != nil {
		t.Fatalf("parse output: %v", err)
	}
	if got := parsed.Query().Get("sslmode"); got != "disable" {
		t.Fatalf("expected sslmode=disable, got %s", got)
	}
}
