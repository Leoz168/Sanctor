package picture

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// StorageClient stores and deletes picture bytes in object storage.
type StorageClient interface {
	Upload(ctx context.Context, key string, content io.Reader, contentType string) (string, error)
	Delete(ctx context.Context, key string) error
}

// SupabaseStorageClient talks to Supabase Storage using its REST API.
type SupabaseStorageClient struct {
	baseURL        string
	serviceRoleKey string
	bucket         string
	httpClient     *http.Client
}

func NewSupabaseStorageClient(baseURL, serviceRoleKey, bucket string) (*SupabaseStorageClient, error) {
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")
	serviceRoleKey = strings.TrimSpace(serviceRoleKey)
	bucket = strings.Trim(strings.TrimSpace(bucket), "/")
	if baseURL == "" {
		return nil, fmt.Errorf("SUPABASE_URL is required")
	}
	if serviceRoleKey == "" {
		return nil, fmt.Errorf("SUPABASE_SERVICE_ROLE_KEY is required")
	}
	if bucket == "" {
		return nil, fmt.Errorf("SUPABASE_STORAGE_BUCKET is required")
	}

	return &SupabaseStorageClient{
		baseURL:        baseURL,
		serviceRoleKey: serviceRoleKey,
		bucket:         bucket,
		httpClient:     http.DefaultClient,
	}, nil
}

func (c *SupabaseStorageClient) Upload(ctx context.Context, key string, content io.Reader, contentType string) (string, error) {
	key = strings.TrimLeft(key, "/")
	endpoint := fmt.Sprintf("%s/storage/v1/object/%s/%s", c.baseURL, url.PathEscape(c.bucket), escapeObjectPath(key))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, content)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.serviceRoleKey)
	req.Header.Set("apikey", c.serviceRoleKey)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("x-upsert", "false")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return "", fmt.Errorf("supabase upload failed with status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	return c.PublicURL(key), nil
}

func (c *SupabaseStorageClient) Delete(ctx context.Context, key string) error {
	key = strings.TrimLeft(key, "/")
	endpoint := fmt.Sprintf("%s/storage/v1/object/%s/%s", c.baseURL, url.PathEscape(c.bucket), escapeObjectPath(key))

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.serviceRoleKey)
	req.Header.Set("apikey", c.serviceRoleKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return fmt.Errorf("supabase delete failed with status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}

func (c *SupabaseStorageClient) PublicURL(key string) string {
	key = strings.TrimLeft(key, "/")
	return fmt.Sprintf("%s/storage/v1/object/public/%s/%s", c.baseURL, url.PathEscape(c.bucket), escapeObjectPath(key))
}

func escapeObjectPath(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	return strings.Join(parts, "/")
}
