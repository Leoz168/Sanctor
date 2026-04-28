package institution

import "github.com/google/uuid"

// Institution represents a post-secondary institution (university, college, etc.)
type Institution struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name    string `json:"name" gorm:"type:varchar(200);not null;uniqueIndex"`
	Country string `json:"country,omitempty" gorm:"type:varchar(100)"`
	Region  string `json:"region,omitempty" gorm:"type:varchar(100)"`
}

// CreateInstitutionRequest represents the data needed to create an institution
type CreateInstitutionRequest struct {
	Name    string `json:"name"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
}

// UpdateInstitutionRequest represents the data that can be updated
type UpdateInstitutionRequest struct {
	Name    string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
}
