package institution

// Repository defines persistence operations for institutions
type Repository interface {
	Create(institution *Institution) error
	FindByID(id string) (*Institution, error)
	FindAll() []*Institution
	Update(institution *Institution) error
	Delete(id string) error
	ExistsByName(name string) bool
}
