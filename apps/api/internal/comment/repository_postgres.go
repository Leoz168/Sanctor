package comment

import (
	"database/sql"
	"errors"

	"sanctor/internal/database"

	"github.com/google/uuid"
)

// PostgresRepository implements Repository for PostgreSQL.
type PostgresRepository struct {
	db *database.DB
}

// NewPostgresRepository creates a new PostgreSQL comment repository.
func NewPostgresRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(comment *Comment) error {
	query := `
		INSERT INTO comments (id, post_id, created_by_user_id, content, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query,
		comment.ID, comment.PostID, comment.CreatedByUserID, comment.Content,
		comment.CreatedAt, comment.UpdatedAt, comment.DeletedAt,
	)
	return err
}

func (r *PostgresRepository) FindByID(id uuid.UUID) (*Comment, error) {
	comment := &Comment{}
	query := `
		SELECT id, post_id, created_by_user_id, content, created_at, updated_at, deleted_at
		FROM comments
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := r.db.QueryRow(query, id).Scan(
		&comment.ID, &comment.PostID, &comment.CreatedByUserID, &comment.Content,
		&comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("comment not found")
	}
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *PostgresRepository) FindByPostID(postID uuid.UUID) []*Comment {
	query := `
		SELECT id, post_id, created_by_user_id, content, created_at, updated_at, deleted_at
		FROM comments
		WHERE post_id = $1 AND deleted_at IS NULL
		ORDER BY created_at ASC
	`

	rows, err := r.db.Query(query, postID)
	if err != nil {
		return []*Comment{}
	}
	defer rows.Close()

	comments := []*Comment{}
	for rows.Next() {
		comment := &Comment{}
		if err := rows.Scan(
			&comment.ID, &comment.PostID, &comment.CreatedByUserID, &comment.Content,
			&comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt,
		); err == nil {
			comments = append(comments, comment)
		}
	}
	return comments
}

func (r *PostgresRepository) Update(comment *Comment) error {
	query := `
		UPDATE comments
		SET content = $2, updated_at = $3
		WHERE id = $1 AND deleted_at IS NULL
	`
	result, err := r.db.Exec(query, comment.ID, comment.Content, comment.UpdatedAt)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("comment not found")
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	query := `
		UPDATE comments
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("comment not found")
	}
	return nil
}

func (r *PostgresRepository) ExistsPost(postID uuid.UUID) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM posts WHERE id = $1)`
	_ = r.db.QueryRow(query, postID).Scan(&exists)
	return exists
}

func (r *PostgresRepository) ExistsUser(userID uuid.UUID) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	_ = r.db.QueryRow(query, userID).Scan(&exists)
	return exists
}
