package post

import (
	"sanctor/internal/database"
	"strings"

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
func (r *GormRepository) Search(filters PostSearchFilters) ([]*Post, error) {
	var posts []*Post
	query := r.db.Model(&Post{})

	if q := strings.TrimSpace(filters.Query); q != "" {
		like := "%" + strings.ToLower(q) + "%"
		query = query.Where(
			"LOWER(title) LIKE ? OR LOWER(content) LIKE ? OR LOWER(description) LIKE ? OR LOWER(address) LIKE ?",
			like,
			like,
			like,
			like,
		)
	}

	if filters.MinPrice != nil {
		query = query.Where("price >= ?", *filters.MinPrice)
	}
	if filters.MaxPrice != nil {
		query = query.Where("price <= ?", *filters.MaxPrice)
	}
	if filters.MinRooms != nil {
		query = query.Where("rooms >= ?", *filters.MinRooms)
	}
	if filters.MinBathrooms != nil {
		query = query.Where("bathrooms >= ?", *filters.MinBathrooms)
	}
	if filters.IsSublet != nil {
		query = query.Where("is_sublet = ?", *filters.IsSublet)
	}
	if filters.Gender != nil {
		query = query.Where("gender = ?", *filters.Gender)
	}
	if filters.Term != nil {
		query = query.Where("term = ?", *filters.Term)
	}
	if pt := strings.TrimSpace(filters.PropertyType); pt != "" {
		query = query.Where("LOWER(property_type) = ?", strings.ToLower(pt))
	}

	if gid := strings.TrimSpace(filters.GroupID); gid != "" {
		query = query.Joins("JOIN post_groups ON post_groups.post_id = posts.id").Where("post_groups.group_id = ?", gid)
	}
	if iid := strings.TrimSpace(filters.InstitutionID); iid != "" {
		query = query.Joins("JOIN post_institutions ON post_institutions.post_id = posts.id").Where("post_institutions.institution_id = ?", iid)
	}

	query = query.Distinct("posts.*")

	sortBy := strings.ToLower(strings.TrimSpace(filters.SortBy))
	sortOrder := strings.ToLower(strings.TrimSpace(filters.SortOrder))
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	sortColumn := "created_at"
	switch sortBy {
	case "price":
		sortColumn = "price"
	case "rooms":
		sortColumn = "rooms"
	case "bathrooms":
		sortColumn = "bathrooms"
	case "updated_at":
		sortColumn = "updated_at"
	}
	query = query.Order(sortColumn + " " + sortOrder)

	if filters.Limit > 0 {
		query = query.Limit(filters.Limit)
	}
	if filters.Offset > 0 {
		query = query.Offset(filters.Offset)
	}

	err := query.Find(&posts).Error
	return posts, err
}
