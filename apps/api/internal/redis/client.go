package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"sanctor/internal/config"

	goredis "github.com/redis/go-redis/v9"
)

// Client wraps the shared Redis client used across modules.
type Client struct {
	inner *goredis.Client
}

// New creates a Redis client from app configuration.
func New(cfg config.RedisConfig) (*Client, error) {
	var options *goredis.Options

	if cfg.SentinelEnabled || strings.TrimSpace(cfg.SentinelAddrs) != "" {
		sentinelAddrs := splitCSV(cfg.SentinelAddrs)
		if len(sentinelAddrs) == 0 {
			return nil, fmt.Errorf("sentinel is enabled but REDIS_SENTINEL_ADDRS is empty")
		}

		failoverOptions := &goredis.FailoverOptions{
			MasterName:       cfg.MasterName,
			SentinelAddrs:    sentinelAddrs,
			Username:         cfg.Username,
			Password:         cfg.Password,
			SentinelUsername: cfg.SentinelUsername,
			SentinelPassword: cfg.SentinelPassword,
			DB:               cfg.DB,
			DialTimeout:      5 * time.Second,
			ReadTimeout:      3 * time.Second,
			WriteTimeout:     3 * time.Second,
		}

		return &Client{inner: goredis.NewFailoverClient(failoverOptions)}, nil
	}

	if cfg.URL != "" {
		opts, err := goredis.ParseURL(cfg.URL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse REDIS_URL: %w", err)
		}
		options = opts
	} else {
		options = &goredis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Password: cfg.Password,
			DB:       cfg.DB,
		}
	}

	if options.DialTimeout == 0 {
		options.DialTimeout = 5 * time.Second
	}
	if options.ReadTimeout == 0 {
		options.ReadTimeout = 3 * time.Second
	}
	if options.WriteTimeout == 0 {
		options.WriteTimeout = 3 * time.Second
	}

	return &Client{inner: goredis.NewClient(options)}, nil
}

// for parsing sentinel addresses
func splitCSV(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed == "" {
			continue
		}
		out = append(out, trimmed)
	}

	return out
}

// Ping verifies connectivity to Redis.
func (c *Client) Ping(ctx context.Context) error {
	if c == nil || c.inner == nil {
		return fmt.Errorf("redis client is not initialized")
	}

	if _, err := c.inner.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}

	return nil
}

// Close closes the underlying Redis client.
func (c *Client) Close() error {
	if c == nil || c.inner == nil {
		return nil
	}
	return c.inner.Close()
}

// Raw exposes the underlying go-redis client for repositories and services.
func (c *Client) Raw() *goredis.Client {
	if c == nil {
		return nil
	}
	return c.inner
}
