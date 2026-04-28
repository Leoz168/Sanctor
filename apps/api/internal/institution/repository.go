package institution

import (
	"errors"

	"github.com/google/uuid"
)

// InMemoryRepository handles data persistence for institutions in memory
type InMemoryRepository struct {
	institutions map[string]*Institution
}

// NewRepository creates a new in-memory institution repository
func NewRepository() Repository {
	return &InMemoryRepository{institutions: make(map[string]*Institution)}
}

func (r *InMemoryRepository) Create(institution *Institution) error {
	if institution == nil {
		return errors.New("institution cannot be nil")
	}
	r.institutions[institution.ID.String()] = institution
	return nil
}

func (r *InMemoryRepository) FindByID(id uuid.UUID) (*Institution, error) {
	inst, exists := r.institutions[id.String()]
	if !exists {
		return nil, errors.New("institution not found")
	}
	return inst, nil
}

func (r *InMemoryRepository) FindAll() []*Institution {
	list := make([]*Institution, 0, len(r.institutions))
	for _, inst := range r.institutions {
		list = append(list, inst)
	}
	return list
}

func (r *InMemoryRepository) Update(institution *Institution) error {
	if institution == nil {
		return errors.New("institution cannot be nil")
	}
	if _, exists := r.institutions[institution.ID.String()]; !exists {
		return errors.New("institution not found")
	}
	r.institutions[institution.ID.String()] = institution
	return nil
}

func (r *InMemoryRepository) Delete(id uuid.UUID) error {
	if _, exists := r.institutions[id.String()]; !exists {
		return errors.New("institution not found")
	}
	delete(r.institutions, id.String())
	return nil
}

func (r *InMemoryRepository) ExistsByName(name string) bool {
	for _, inst := range r.institutions {
		if inst.Name == name {
			return true
		}
	}
	return false
}
