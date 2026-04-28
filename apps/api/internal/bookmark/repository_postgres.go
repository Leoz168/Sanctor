package bookmark

import (
	"sanctor/internal/database"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

// PostgresRepository implements Repository for PostgreSQL.
type PostgresRepository struct {
	db *database.DB
}

// NewPostgresRepository creates a new PostgreSQL bookmark repository.
func NewPostgresRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(bookmark *Bookmark) error {
	// Composite primary key enforces uniqueness; this avoids duplicate insert errors.
	return r.db.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(bookmark).Error
}

func (r *PostgresRepository) Delete(userID, postID uuid.UUID) error {
	return r.db.Gorm.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&Bookmark{}).Error
}

func (r *PostgresRepository) FindByUserID(userID uuid.UUID) ([]*Bookmark, error) {
	bookmarks := make([]*Bookmark, 0)
	err := r.db.Gorm.Where("user_id = ?", userID).Order("created_at DESC").Find(&bookmarks).Error
	return bookmarks, err
}

func (r *PostgresRepository) FindPostsByUserID(userID uuid.UUID) ([]*BookmarkedPost, error) {
	posts := make([]*BookmarkedPost, 0)
	err := r.db.Gorm.
		Table("posts").
		Select("posts.*").
		Joins("JOIN post_bookmarks ON post_bookmarks.post_id = posts.id").
		Where("post_bookmarks.user_id = ?", userID).
		Order("post_bookmarks.created_at DESC").
		Find(&posts).Error

	return posts, err
}

func (r *PostgresRepository) Exists(userID, postID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Gorm.Model(&Bookmark{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresRepository) ExistsPost(postID uuid.UUID) bool {
	var count int64
	err := r.db.Gorm.Table("posts").Where("id = ?", postID).Count(&count).Error
	return err == nil && count > 0
}

func (r *PostgresRepository) ExistsUser(userID uuid.UUID) bool {
	var count int64
	err := r.db.Gorm.Table("users").Where("id = ?", userID).Count(&count).Error
	return err == nil && count > 0
}

var _ Repository = (*PostgresRepository)(nil)
