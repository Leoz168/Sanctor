package picture

import (
	"errors"

	"sanctor/internal/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repository defines picture metadata persistence.
type Repository interface {
	Create(picture *Picture) error
	FindByID(id uuid.UUID) (*Picture, error)
	FindByPostID(postID uuid.UUID) ([]*Picture, error)
	Delete(id uuid.UUID) error
	ExistsPost(postID uuid.UUID) bool
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

func (r *GormRepository) FindByPostID(postID uuid.UUID) ([]*Picture, error) {
	var pictures []*Picture
	err := r.db.Where("post_id = ?", postID).Order("display_order ASC, created_at ASC").Find(&pictures).Error
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

func (r *GormRepository) ExistsPost(postID uuid.UUID) bool {
	var count int64
	if err := r.db.Table("posts").Where("id = ?", postID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
