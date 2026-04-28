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

func (r *PostgresRepository) Delete(userID, postID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return err
	}
	return r.db.Gorm.Where("user_id = ? AND post_id = ?", userUUID, postUUID).Delete(&Bookmark{}).Error
}

func (r *PostgresRepository) FindByUserID(userID string) ([]*Bookmark, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	bookmarks := make([]*Bookmark, 0)
	err = r.db.Gorm.Where("user_id = ?", userUUID).Order("created_at DESC").Find(&bookmarks).Error
	return bookmarks, err
}

func (r *PostgresRepository) FindPostsByUserID(userID string) ([]*BookmarkedPost, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	posts := make([]*BookmarkedPost, 0)
	err = r.db.Gorm.
		Table("posts").
		Select("posts.*").
		Joins("JOIN post_bookmarks ON post_bookmarks.post_id = posts.id").
		Where("post_bookmarks.user_id = ?", userUUID).
		Order("post_bookmarks.created_at DESC").
		Find(&posts).Error

	return posts, err
}

func (r *PostgresRepository) Exists(userID, postID string) (bool, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return false, err
	}
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return false, err
	}
	var count int64
	err = r.db.Gorm.Model(&Bookmark{}).Where("user_id = ? AND post_id = ?", userUUID, postUUID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresRepository) ExistsPost(postID string) bool {
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return false
	}
	var count int64
	err = r.db.Gorm.Table("posts").Where("id = ?", postUUID).Count(&count).Error
	return err == nil && count > 0
}

func (r *PostgresRepository) ExistsUser(userID string) bool {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return false
	}
	var count int64
	err = r.db.Gorm.Table("users").Where("id = ?", userUUID).Count(&count).Error
	return err == nil && count > 0
}

var _ Repository = (*PostgresRepository)(nil)
