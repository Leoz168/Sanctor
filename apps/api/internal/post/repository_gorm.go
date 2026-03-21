package post

import (
	"sanctor/internal/database"

	"gorm.io/gorm"
)

// GormRepository handles data persistence for posts using GORM
type GormRepository struct {
	db *gorm.DB
}

// NewGormRepository creates a new GORM post repository
func NewGormRepository(db *database.DB) *GormRepository {
	return &GormRepository{
		db: db.Gorm,
	}
}

// Create adds a new post
func (r *GormRepository) CreateWithLinks(post *Post, groupIDs []string, institutionIDs []string) (*Post, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(post).Error; err != nil {
			return err
		}

		if len(groupIDs) > 0 {
			postGroups := make([]PostGroup, 0, len(groupIDs))
			for _, groupID := range groupIDs {
				postGroups = append(postGroups, PostGroup{
					PostID:   post.ID,
					GroupID:  groupID,
					LinkedAt: post.CreatedAt,
				})
			}

			if err := tx.Create(&postGroups).Error; err != nil {
				return err
			}
		}

		if len(institutionIDs) > 0 {
			postInstitutions := make([]PostInstitution, 0, len(institutionIDs))
			for _, institutionID := range institutionIDs {
				postInstitutions = append(postInstitutions, PostInstitution{
					PostID:        post.ID,
					InstitutionID: institutionID,
					LinkedAt:      post.CreatedAt,
				})
			}

			if err := tx.Create(&postInstitutions).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

// FindByID retrieves a post by ID
func (r *GormRepository) FindByID(id string) (*Post, error) {
	var post Post
	err := r.db.Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// FindAll retrieves all posts
func (r *GormRepository) FindAll() ([]*Post, error) {
	var posts []*Post
	err := r.db.Find(&posts).Error
	return posts, err
}

// FindByUserID retrieves all posts for a specific user
func (r *GormRepository) FindByUserID(userID string) ([]*Post, error) {
	var posts []*Post
	err := r.db.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// Update updates a post
func (r *GormRepository) Update(post *Post) error {
	return r.db.Save(post).Error
}

// Delete removes a post
func (r *GormRepository) Delete(id string) error {
	return r.db.Delete(&Post{}, "id = ?", id).Error
}

// Search posts by filters
func (r *GormRepository) Search(filters map[string]interface{}) ([]*Post, error) {
	var posts []*Post
	query := r.db

	// Apply filters
	if term, ok := filters["term"].(Term); ok {
		query = query.Where("term = ?", term)
	}
	if gender, ok := filters["gender"].(string); ok && gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if propertyType, ok := filters["propertyType"].(string); ok && propertyType != "" {
		query = query.Where("property_type = ?", propertyType)
	}
	if maxPrice, ok := filters["maxPrice"].(string); ok && maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	err := query.Find(&posts).Error
	return posts, err
}
