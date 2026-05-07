package picture

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

const MaxUploadBytes = 10 << 20

var allowedImageTypes = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
	"image/gif":  ".gif",
}

// Service handles picture upload and metadata operations.
type Service struct {
	repo    Repository
	storage StorageClient
}

func NewService(repo Repository, storage StorageClient) *Service {
	return &Service{repo: repo, storage: storage}
}

func (s *Service) UploadPicture(ctx context.Context, postID uuid.UUID, file io.Reader, contentType string, originalFilename string, caption string, order int) (*Picture, error) {
	if s.repo == nil {
		return nil, errors.New("picture repository not initialized")
	}
	if s.storage == nil {
		return nil, errors.New("picture storage not initialized")
	}
	if postID == uuid.Nil {
		return nil, errors.New("postId is required")
	}
	if !s.repo.ExistsPost(postID) {
		return nil, errors.New("post not found")
	}

	contentType = normalizeContentType(contentType)
	extension, ok := allowedImageTypes[contentType]
	if !ok {
		return nil, errors.New("unsupported image type")
	}
	if ext := strings.ToLower(filepath.Ext(originalFilename)); ext != "" && extensionForContentType(contentType, ext) {
		extension = ext
	}

	pictureID := uuid.New()
	storageKey := fmt.Sprintf("posts/%s/%s%s", postID.String(), pictureID.String(), extension)

	url, err := s.storage.Upload(ctx, storageKey, file, contentType)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	picture := &Picture{
		ID:         pictureID,
		PostID:     postID,
		URL:        url,
		StorageKey: storageKey,
		Caption:    strings.TrimSpace(caption),
		Order:      order,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if err := s.repo.Create(picture); err != nil {
		_ = s.storage.Delete(ctx, storageKey)
		return nil, err
	}
	return picture, nil
}

func (s *Service) GetPicturesByPost(postID uuid.UUID) ([]*Picture, error) {
	if s.repo == nil {
		return nil, errors.New("picture repository not initialized")
	}
	if postID == uuid.Nil {
		return nil, errors.New("postId is required")
	}
	return s.repo.FindByPostID(postID)
}

func (s *Service) DeletePicture(ctx context.Context, id uuid.UUID) error {
	if s.repo == nil {
		return errors.New("picture repository not initialized")
	}
	if id == uuid.Nil {
		return errors.New("picture id is required")
	}

	picture, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if s.storage != nil {
		if err := s.storage.Delete(ctx, picture.StorageKey); err != nil {
			return err
		}
	}
	return s.repo.Delete(id)
}

func normalizeContentType(value string) string {
	contentType := strings.TrimSpace(strings.ToLower(value))
	if mediaType, _, err := mime.ParseMediaType(contentType); err == nil {
		contentType = mediaType
	}
	return contentType
}

func extensionForContentType(contentType string, extension string) bool {
	switch contentType {
	case "image/jpeg":
		return extension == ".jpg" || extension == ".jpeg"
	case "image/png":
		return extension == ".png"
	case "image/webp":
		return extension == ".webp"
	case "image/gif":
		return extension == ".gif"
	default:
		return false
	}
}
