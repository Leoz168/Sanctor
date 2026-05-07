package picture

import (
	"errors"

	"sanctor/internal/database"
	sharedtypes "sanctor/pkg/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repository defines picture metadata persistence.
type Repository interface {
	Create(picture *Picture) error
	FindByID(id uuid.UUID) (*Picture, error)
	FindByOwner(ownerType sharedtypes.OwnerType, ownerID uuid.UUID) ([]*Picture, error)
	Delete(id uuid.UUID) error
	ExistsOwner(ownerType sharedtypes.OwnerType, ownerID uuid.UUID) bool
}

// GormRepository stores picture metadata with GORM.
type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *database.DB) Repository {
	return &GormRepository{db: db.Gorm}
}

func (r *GormRepository) Create(picture *Picture) error {
	return r.db.Create(picture).Error
}

func (r *GormRepository) FindByID(id uuid.UUID) (*Picture, error) {
	var picture Picture
	err := r.db.Where("id = ?", id).First(&picture).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("picture not found")
	}
	if err != nil {
		return nil, err
	}
	return &picture, nil
}

func (r *GormRepository) FindByOwner(ownerType sharedtypes.OwnerType, ownerID uuid.UUID) ([]*Picture, error) {
	var pictures []*Picture
	err := r.db.Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).Order("display_order ASC, created_at ASC").Find(&pictures).Error
	return pictures, err
}

func (r *GormRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&Picture{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("picture not found")
	}
	return nil
}

func (r *GormRepository) ExistsOwner(ownerType sharedtypes.OwnerType, ownerID uuid.UUID) bool {
	var count int64
	tableName := ""
	switch ownerType {
	case sharedtypes.OwnerTypePost:
		tableName = "posts"
	case sharedtypes.OwnerTypeCommunity:
		tableName = "communities"
	default:
		return false
	}
	if err := r.db.Table(tableName).Where("id = ?", ownerID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
