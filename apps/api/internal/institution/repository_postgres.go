package institution

import (
	"database/sql"
	"errors"

	"sanctor/internal/database"

	"github.com/google/uuid"
)

// PostgresRepository implements Repository for PostgreSQL
type PostgresRepository struct {
	db *database.DB
}

// NewPostgresRepository creates a new PostgreSQL institution repository
func NewPostgresRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(institution *Institution) error {
	query := `
		INSERT INTO institutions (id, name, country, region)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(query, institution.ID, institution.Name, institution.Country, institution.Region)
	return err
}

func (r *PostgresRepository) FindByID(id uuid.UUID) (*Institution, error) {
	inst := &Institution{}
	query := `SELECT id, name, country, region FROM institutions WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&inst.ID, &inst.Name, &inst.Country, &inst.Region)
	if err == sql.ErrNoRows {
		return nil, errors.New("institution not found")
	}
	return inst, err
}

func (r *PostgresRepository) FindAll() []*Institution {
	query := `SELECT id, name, country, region FROM institutions ORDER BY name ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return []*Institution{}
	}
	defer rows.Close()

	institutions := []*Institution{}
	for rows.Next() {
		inst := &Institution{}
		if err := rows.Scan(&inst.ID, &inst.Name, &inst.Country, &inst.Region); err == nil {
			institutions = append(institutions, inst)
		}
	}
	return institutions
}

func (r *PostgresRepository) Update(institution *Institution) error {
	query := `UPDATE institutions SET name = $2, country = $3, region = $4 WHERE id = $1`
	result, err := r.db.Exec(query, institution.ID, institution.Name, institution.Country, institution.Region)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("institution not found")
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM institutions WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("institution not found")
	}
	return nil
}

func (r *PostgresRepository) ExistsByName(name string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM institutions WHERE name = $1)`
	_ = r.db.QueryRow(query, name).Scan(&exists)
	return exists
}
