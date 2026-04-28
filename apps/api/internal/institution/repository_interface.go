package institution

import "github.com/google/uuid"

// Repository defines persistence operations for institutions
type Repository interface {
	Create(institution *Institution) error
	FindByID(id uuid.UUID) (*Institution, error)
	FindAll() []*Institution
	Update(institution *Institution) error
	Delete(id uuid.UUID) error
	ExistsByName(name string) bool
}
