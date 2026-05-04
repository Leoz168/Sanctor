package redis

import (
	"testing"

	"sanctor/internal/config"
)

func TestSplitCSV(t *testing.T) {
	out := splitCSV("a, b,,c")
	if len(out) != 3 {
		t.Fatalf("expected 3 values, got %v", out)
	}
	if out[0] != "a" || out[1] != "b" || out[2] != "c" {
		t.Fatalf("unexpected values: %v", out)
	}

	out = splitCSV(" ")
	if out != nil {
		t.Fatalf("expected nil for empty input")
	}
}

func TestNewSentinelRequiresAddrs(t *testing.T) {
	_, err := New(config.RedisConfig{SentinelEnabled: true})
	if err == nil {
		t.Fatalf("expected error for missing sentinel addrs")
	}
}

func TestNewInvalidURL(t *testing.T) {
	_, err := New(config.RedisConfig{URL: "http://localhost"})
	if err == nil {
		t.Fatalf("expected error for invalid redis url scheme")
	}
}

func TestRawNil(t *testing.T) {
	var c *Client
	if c.Raw() != nil {
		t.Fatalf("expected nil raw client")
	}
}
