package post

import (
	"fmt"
	"sort"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repository handles data persistence for posts
type Repository struct {
	posts map[string]*Post
	db    *gorm.DB
}

// NewRepository creates a new post repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		posts: make(map[string]*Post),
		db:    db,
	}
}

// CreateWithLinks adds a new post and ignores relation links in memory mode
func (r *Repository) CreateWithLinks(post *Post, groupIDs []uuid.UUID, institutionIDs []uuid.UUID) (*Post, error) {
	r.posts[post.ID.String()] = post
	return post, nil
}

// FindByID retrieves a post by ID from the database
func (r *Repository) FindByID(id uuid.UUID) (*Post, error) {
	var post Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("post not found")
		}
		return nil, err
	}
	return &post, nil
}

// FindAll retrieves all posts
func (r *Repository) FindAll() ([]*Post, error) {
	posts := make([]*Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

// Search retrieves posts in memory using basic filtering, sorting, and pagination.
func (r *Repository) Search(filters PostSearchFilters) ([]*Post, error) {
	posts := make([]*Post, 0, len(r.posts))
	q := strings.ToLower(strings.TrimSpace(filters.Query))

	for _, p := range r.posts {
		if q != "" {
			haystack := strings.ToLower(p.Title + " " + p.Content + " " + p.Description + " " + p.Address)
			if !strings.Contains(haystack, q) {
				continue
			}
		}

		if filters.MinPrice != nil && p.Price < *filters.MinPrice {
			continue
		}
		if filters.MaxPrice != nil && p.Price > *filters.MaxPrice {
			continue
		}
		if filters.MinRooms != nil && p.Rooms < *filters.MinRooms {
			continue
		}
		if filters.MinBathrooms != nil && p.Bathrooms < *filters.MinBathrooms {
			continue
		}
		if filters.IsSublet != nil && p.IsSublet != *filters.IsSublet {
			continue
		}
		if filters.Gender != nil && p.Gender != *filters.Gender {
			continue
		}
		if filters.Term != nil && p.Term != *filters.Term {
			continue
		}
		if filters.PropertyType != "" && !strings.EqualFold(p.PropertyType, filters.PropertyType) {
			continue
		}

		posts = append(posts, p)
	}

	sortBy := strings.ToLower(strings.TrimSpace(filters.SortBy))
	sortOrder := strings.ToLower(strings.TrimSpace(filters.SortOrder))
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	sort.Slice(posts, func(i, j int) bool {
		left := posts[i]
		right := posts[j]

		less := false
		switch sortBy {
		case "price":
			less = left.Price < right.Price
		case "rooms":
			less = left.Rooms < right.Rooms
		case "bathrooms":
			less = left.Bathrooms < right.Bathrooms
		case "updated_at":
			less = left.UpdatedAt.Before(right.UpdatedAt)
		default:
			less = left.CreatedAt.Before(right.CreatedAt)
		}

		if sortOrder == "asc" {
			return less
		}
		return !less
	})

	start := filters.Offset
	if start < 0 {
		start = 0
	}
	if start > len(posts) {
		start = len(posts)
	}

	end := len(posts)
	if filters.Limit > 0 {
		end = start + filters.Limit
		if end > len(posts) {
			end = len(posts)
		}
	}

	return posts[start:end], nil
}

// Update updates a post
func (r *Repository) Update(post *Post) error {
	r.posts[post.ID.String()] = post
	return nil
}

// Delete removes a post
func (r *Repository) Delete(id uuid.UUID) error {
	delete(r.posts, id.String())
	return nil
}
