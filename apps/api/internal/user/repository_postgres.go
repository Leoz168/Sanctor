package user

import (
	"database/sql"
	"errors"

	"sanctor/internal/database"

	"github.com/google/uuid"
)

// PostgresRepository implements Repository interface for PostgreSQL
type PostgresRepository struct {
	db *database.DB
}

// NewPostgresRepository creates a new PostgreSQL user repository
func NewPostgresRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}

// Create adds a new user to the database
func (r *PostgresRepository) Create(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	query := `
		INSERT INTO users (
			id, email, username, password_hash, google_sub,
			avatar, bio, is_active, is_verified, is_blacklisted, last_login_at,
			created_at, updated_at, gender, age, institution_id, major
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`

	_, err := r.db.Exec(query,
		user.ID, user.Email, user.Username, user.PasswordHash, user.GoogleSub,
		user.Avatar, user.Bio, user.IsActive, user.IsVerified, user.IsBlacklisted,
		user.LastLoginAt, user.CreatedAt, user.UpdatedAt, user.Gender,
		user.Age, user.InstitutionID, user.Major,
	)

	return err
}

// FindByID retrieves a user by ID
func (r *PostgresRepository) FindByID(id uuid.UUID) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, username, password_hash, google_sub,
		       avatar, bio, is_active, is_verified, is_blacklisted, last_login_at,
		       created_at, updated_at, gender, age, institution_id, major
		FROM users WHERE id = $1
	`

	var googleSub sql.NullString
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Username, &user.PasswordHash, &googleSub,
		&user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
		&user.IsBlacklisted, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		&user.Gender, &user.Age, &user.InstitutionID, &user.Major,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	if googleSub.Valid {
		user.GoogleSub = &googleSub.String
	}

	return user, nil
}

// FindAll retrieves all users
func (r *PostgresRepository) FindAll() []*User {
	query := `
		SELECT id, email, username, password_hash,
		       google_sub, avatar, bio, is_active, is_verified, is_blacklisted, last_login_at,
		       created_at, updated_at, gender, age, institution_id, major
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return []*User{}
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		user := &User{}
		var googleSub sql.NullString
		err := rows.Scan(
			&user.ID, &user.Email, &user.Username, &user.PasswordHash, &googleSub,
			&user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
			&user.IsBlacklisted, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
			&user.Gender, &user.Age, &user.InstitutionID, &user.Major,
		)
		if err == nil {
			if googleSub.Valid {
				user.GoogleSub = &googleSub.String
			}
			users = append(users, user)
		}
	}

	return users
}

// Update modifies an existing user
func (r *PostgresRepository) Update(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	query := `
		UPDATE users SET
			email = $2, username = $3, password_hash = $4, google_sub = $5,
			avatar = $6, bio = $7, is_active = $8, is_verified = $9,
			is_blacklisted = $10, last_login_at = $11, updated_at = $12,
			gender = $13, age = $14, institution_id = $15, major = $16
		WHERE id = $1
	`

	result, err := r.db.Exec(query,
		user.ID, user.Email, user.Username, user.PasswordHash, user.GoogleSub,
		user.Avatar, user.Bio, user.IsActive, user.IsVerified,
		user.IsBlacklisted, user.LastLoginAt, user.UpdatedAt, user.Gender,
		user.Age, user.InstitutionID, user.Major,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

// Delete removes a user from the database
func (r *PostgresRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

// ExistsByEmail checks if a user with the given email exists
func (r *PostgresRepository) ExistsByEmail(email string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := r.db.QueryRow(query, email).Scan(&exists)
	return err == nil && exists
}

// IsEmailBlacklisted checks if a user with the given email is blacklisted.
func (r *PostgresRepository) IsEmailBlacklisted(email string) (bool, error) {
	var isBlacklisted bool
	query := `SELECT is_blacklisted FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&isBlacklisted)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return isBlacklisted, nil
}

// ExistsByUsername checks if a user with the given username exists
func (r *PostgresRepository) ExistsByUsername(username string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := r.db.QueryRow(query, username).Scan(&exists)
	return err == nil && exists
}

// FindByEmail retrieves a user by email
func (r *PostgresRepository) FindByEmail(email string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, username, password_hash, google_sub,
		       avatar, bio, is_active, is_verified, is_blacklisted, last_login_at,
		       created_at, updated_at, gender, age, institution_id, major
		FROM users WHERE email = $1
	`

	var googleSub sql.NullString
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Username, &user.PasswordHash, &googleSub,
		&user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
		&user.IsBlacklisted, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		&user.Gender, &user.Age, &user.InstitutionID, &user.Major,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	if googleSub.Valid {
		user.GoogleSub = &googleSub.String
	}

	return user, nil
}

// FindByUsername retrieves a user by username
func (r *PostgresRepository) FindByUsername(username string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, username, password_hash, google_sub,
		       avatar, bio, is_active, is_verified, is_blacklisted, last_login_at,
		       created_at, updated_at, gender, age, institution_id, major
		FROM users WHERE username = $1
	`

	var googleSub sql.NullString
	err := r.db.QueryRow(query, username).Scan(
		&user.ID, &user.Email, &user.Username, &user.PasswordHash, &googleSub,
		&user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
		&user.IsBlacklisted, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		&user.Gender, &user.Age, &user.InstitutionID, &user.Major,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	if googleSub.Valid {
		user.GoogleSub = &googleSub.String
	}

	return user, nil
}

// FindByGoogleSub retrieves a user by Google subject.
func (r *PostgresRepository) FindByGoogleSub(sub string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, username, password_hash, google_sub,
		       avatar, bio, is_active, is_verified, is_blacklisted, last_login_at,
		       created_at, updated_at, gender, age, institution_id, major
		FROM users WHERE google_sub = $1
	`

	var googleSub sql.NullString
	err := r.db.QueryRow(query, sub).Scan(
		&user.ID, &user.Email, &user.Username, &user.PasswordHash, &googleSub,
		&user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
		&user.IsBlacklisted, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		&user.Gender, &user.Age, &user.InstitutionID, &user.Major,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	if googleSub.Valid {
		user.GoogleSub = &googleSub.String
	}

	return user, nil
}
