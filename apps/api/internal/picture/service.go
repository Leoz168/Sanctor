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

	sharedtypes "sanctor/pkg/types"

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

func (s *Service) UploadPicture(ctx context.Context, ownerType sharedtypes.OwnerType, ownerID uuid.UUID, file io.Reader, contentType string, originalFilename string, caption string, order int) (*Picture, error) {
	if s.repo == nil {
		return nil, errors.New("picture repository not initialized")
	}
	if s.storage == nil {
		return nil, errors.New("picture storage not initialized")
	}
	if err := validateOwner(ownerType, ownerID); err != nil {
		return nil, err
	}
	if !s.repo.ExistsOwner(ownerType, ownerID) {
		return nil, fmt.Errorf("%s not found", ownerType)
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
	storageKey := fmt.Sprintf("%ss/%s/%s%s", ownerType, ownerID.String(), pictureID.String(), extension)

	url, err := s.storage.Upload(ctx, storageKey, file, contentType)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	picture := &Picture{
		ID:         pictureID,
		OwnerType:  ownerType,
		OwnerID:    ownerID,
		URL:        url,
		StorageKey: storageKey,
		Caption:    strings.TrimSpace(caption),
		Order:      order,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if ownerType == sharedtypes.OwnerTypePost {
		picture.PostID = &ownerID
	}
	if err := s.repo.Create(picture); err != nil {
		_ = s.storage.Delete(ctx, storageKey)
		return nil, err
	}
	return picture, nil
}

func (s *Service) GetPicturesByOwner(ownerType sharedtypes.OwnerType, ownerID uuid.UUID) ([]*Picture, error) {
	if s.repo == nil {
		return nil, errors.New("picture repository not initialized")
	}
	if err := validateOwner(ownerType, ownerID); err != nil {
		return nil, err
	}
	return s.repo.FindByOwner(ownerType, ownerID)
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

func validateOwner(ownerType sharedtypes.OwnerType, ownerID uuid.UUID) error {
	if ownerID == uuid.Nil {
		return errors.New("ownerId is required")
	}
	switch ownerType {
	case sharedtypes.OwnerTypePost, sharedtypes.OwnerTypeCommunity:
		return nil
	default:
		return errors.New("ownerType must be post or community")
	}
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
