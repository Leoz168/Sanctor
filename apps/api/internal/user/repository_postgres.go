package user

import (
	"database/sql"
	"errors"

	"sanctor/internal/database"
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
			id, email, username, first_name, last_name, password_hash,
			avatar, bio, is_active, is_verified,last_login_at,
			created_at, updated_at, gender, age, university, major
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`

	_, err := r.db.Exec(query,
		user.ID, user.Email, user.Username, user.FirstName, user.LastName,
		user.PasswordHash, user.Avatar, user.Bio, user.IsActive, user.IsVerified,
		user.LastLoginAt, user.CreatedAt, user.UpdatedAt,
		user.Gender, user.Age, user.University, user.Major,
	)

	return err
}

// FindByID retrieves a user by ID
func (r *PostgresRepository) FindByID(id string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, username, first_name, last_name, password_hash,
		       avatar, bio, is_active, is_verified, last_login_at,
		       created_at, updated_at, gender, age, university, major
		FROM users WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName,
		&user.PasswordHash, &user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
		&user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		&user.Gender, &user.Age, &user.University, &user.Major,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindAll retrieves all users
func (r *PostgresRepository) FindAll() []*User {
	query := `
		SELECT id, email, username, first_name, last_name, password_hash,
		       avatar, bio, is_active, is_verified, last_login_at,
		       created_at, updated_at, gender, age, university, major
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
		err := rows.Scan(
			&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName,
			&user.PasswordHash, &user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
			&user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
			&user.Gender, &user.Age, &user.University, &user.Major,
		)
		if err == nil {
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
			email = $2, username = $3, first_name = $4, last_name = $5,
			password_hash = $6, avatar = $7, bio = $8, is_active = $9,
			is_verified = $10, last_login_at = $11, updated_at = $12,
			gender = $13, age = $14, university = $15, major = $16
		WHERE id = $1
	`

	result, err := r.db.Exec(query,
		user.ID, user.Email, user.Username, user.FirstName, user.LastName,
		user.PasswordHash, user.Avatar, user.Bio, user.IsActive, user.IsVerified,
		user.LastLoginAt, user.UpdatedAt,
		user.Gender, user.Age, user.University, user.Major,
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
func (r *PostgresRepository) Delete(id string) error {
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
		SELECT id, email, username, first_name, last_name, password_hash,
		       avatar, bio, is_active, is_verified, last_login_at,
		       created_at, updated_at, gender, age, university, major
		FROM users WHERE email = $1
	`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName,
		&user.PasswordHash, &user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
		&user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		&user.Gender, &user.Age, &user.University, &user.Major,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByUsername retrieves a user by username
func (r *PostgresRepository) FindByUsername(username string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, username, first_name, last_name, password_hash,
		       avatar, bio, is_active, is_verified, last_login_at,
		       created_at, updated_at, gender, age, university, major
		FROM users WHERE username = $1
	`

	err := r.db.QueryRow(query, username).Scan(
		&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName,
		&user.PasswordHash, &user.Avatar, &user.Bio, &user.IsActive, &user.IsVerified,
		&user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		&user.Gender, &user.Age, &user.University, &user.Major,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
