package blocking

import (
	"sanctor/internal/database"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type PostgresRepository struct {
	db *database.DB
}

func NewPostgresRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(block *UserBlock) error {
	return r.db.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(block).Error
}

func (r *PostgresRepository) Delete(blockerID, blockeeID uuid.UUID) error {
	return r.db.Gorm.Where("blocker = ? AND blockee = ?", blockerID, blockeeID).Delete(&UserBlock{}).Error
}

func (r *PostgresRepository) FindByBlockerID(blockerID uuid.UUID) ([]*UserBlock, error) {
	blocks := make([]*UserBlock, 0)
	err := r.db.Gorm.Where("blocker = ?", blockerID).Order("linked_at DESC").Find(&blocks).Error
	return blocks, err
}

func (r *PostgresRepository) FindByBlockeeID(blockeeID uuid.UUID) ([]*UserBlock, error) {
	blocks := make([]*UserBlock, 0)
	err := r.db.Gorm.Where("blockee = ?", blockeeID).Order("linked_at DESC").Find(&blocks).Error
	return blocks, err
}

func (r *PostgresRepository) Exists(blockerID, blockeeID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Gorm.Model(&UserBlock{}).Where("blocker = ? AND blockee = ?", blockerID, blockeeID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresRepository) ExistsUser(userID uuid.UUID) bool {
	var count int64
	err := r.db.Gorm.Table("users").Where("id = ?", userID).Count(&count).Error
	return err == nil && count > 0
}
