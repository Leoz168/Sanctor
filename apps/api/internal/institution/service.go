package institution

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Service handles business logic for institution operations
type Service struct {
	repo Repository
}

// NewService creates a new institution service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateInstitution(req CreateInstitutionRequest) (*Institution, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("institution name is required")
	}
	if s.repo.ExistsByName(name) {
		return nil, errors.New("institution already exists")
	}

	institution := &Institution{
		ID:      uuid.New(),
		Name:    name,
		Country: strings.TrimSpace(req.Country),
		Region:  strings.TrimSpace(req.Region),
	}

	if err := s.repo.Create(institution); err != nil {
		return nil, err
	}
	return institution, nil
}

func (s *Service) GetInstitution(id string) (*Institution, error) {
	if id == "" {
		return nil, errors.New("institution ID is required")
	}
	return s.repo.FindByID(id)
}

func (s *Service) GetAllInstitutions() ([]*Institution, error) {
	return s.repo.FindAll(), nil
}

func (s *Service) UpdateInstitution(id string, req UpdateInstitutionRequest) (*Institution, error) {
	institution, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("institution not found")
	}

	if req.Name != "" {
		name := strings.TrimSpace(req.Name)
		if name == "" {
			return nil, errors.New("institution name cannot be empty")
		}
		if name != institution.Name && s.repo.ExistsByName(name) {
			return nil, errors.New("institution already exists")
		}
		institution.Name = name
	}
	if req.Country != "" {
		institution.Country = strings.TrimSpace(req.Country)
	}
	if req.Region != "" {
		institution.Region = strings.TrimSpace(req.Region)
	}

	if err := s.repo.Update(institution); err != nil {
		return nil, err
	}
	return institution, nil
}

func (s *Service) DeleteInstitution(id string) error {
	if id == "" {
		return errors.New("institution ID is required")
	}
	return s.repo.Delete(id)
}
