package group

import (
	"database/sql"
	"errors"

	"sanctor/internal/database"
)

// PostgresRepository implements Repository interface for PostgreSQL
type PostgresRepository struct {
	db *database.DB
}

// NewPostgresRepository creates a new PostgreSQL group repository
func NewPostgresRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}

// Create creates a new group
func (r *PostgresRepository) Create(group *Group) error {
	query := `
		INSERT INTO groups (id, name, description, is_private, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query, group.ID, group.Name, group.Description, group.IsPrivate,
		group.CreatedBy, group.CreatedAt, group.UpdatedAt)
	return err
}

// FindByID finds a group by ID
func (r *PostgresRepository) FindByID(id string) (*Group, error) {
	group := &Group{}
	query := `SELECT id, name, description, is_private, created_by, created_at, updated_at 
	          FROM groups WHERE id = $1`
	
	err := r.db.QueryRow(query, id).Scan(&group.ID, &group.Name, &group.Description,
		&group.IsPrivate, &group.CreatedBy, &group.CreatedAt, &group.UpdatedAt)
	
	if err == sql.ErrNoRows {
		return nil, errors.New("group not found")
	}
	return group, err
}

// FindAll returns all groups
func (r *PostgresRepository) FindAll() []*Group {
	query := `SELECT id, name, description, is_private, created_by, created_at, updated_at 
	          FROM groups ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return []*Group{}
	}
	defer rows.Close()

	groups := []*Group{}
	for rows.Next() {
		group := &Group{}
		if err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.IsPrivate,
			&group.CreatedBy, &group.CreatedAt, &group.UpdatedAt); err == nil {
			groups = append(groups, group)
		}
	}
	return groups
}

// Update updates an existing group
func (r *PostgresRepository) Update(group *Group) error {
	query := `UPDATE groups SET name = $2, description = $3, is_private = $4, updated_at = $5 
	          WHERE id = $1`
	
	result, err := r.db.Exec(query, group.ID, group.Name, group.Description, 
		group.IsPrivate, group.UpdatedAt)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("group not found")
	}
	return nil
}

// Delete deletes a group (CASCADE will delete user_groups automatically)
func (r *PostgresRepository) Delete(id string) error {
	query := `DELETE FROM groups WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("group not found")
	}
	return nil
}

// AddUserToGroup adds a user to a group
func (r *PostgresRepository) AddUserToGroup(userGroup *UserGroup) error {
	query := `INSERT INTO user_groups (user_id, group_id, role, joined_at) 
	          VALUES ($1, $2, $3, $4)`
	
	_, err := r.db.Exec(query, userGroup.UserID, userGroup.GroupID, 
		userGroup.Role, userGroup.JoinedAt)
	return err
}

// RemoveUserFromGroup removes a user from a group
func (r *PostgresRepository) RemoveUserFromGroup(userID, groupID string) error {
	query := `DELETE FROM user_groups WHERE user_id = $1 AND group_id = $2`
	result, err := r.db.Exec(query, userID, groupID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("user not in group")
	}
	return nil
}

// GetGroupMembers returns all members of a group
func (r *PostgresRepository) GetGroupMembers(groupID string) ([]*UserGroup, error) {
	query := `SELECT user_id, group_id, role, joined_at 
	          FROM user_groups WHERE group_id = $1 ORDER BY joined_at`
	
	rows, err := r.db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := []*UserGroup{}
	for rows.Next() {
		ug := &UserGroup{}
		if err := rows.Scan(&ug.UserID, &ug.GroupID, &ug.Role, &ug.JoinedAt); err == nil {
			members = append(members, ug)
		}
	}
	return members, nil
}

// GetUserGroups returns all groups a user belongs to
func (r *PostgresRepository) GetUserGroups(userID string) []*UserGroup {
	query := `SELECT user_id, group_id, role, joined_at 
	          FROM user_groups WHERE user_id = $1 ORDER BY joined_at DESC`
	
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return []*UserGroup{}
	}
	defer rows.Close()

	userGroups := []*UserGroup{}
	for rows.Next() {
		ug := &UserGroup{}
		if err := rows.Scan(&ug.UserID, &ug.GroupID, &ug.Role, &ug.JoinedAt); err == nil {
			userGroups = append(userGroups, ug)
		}
	}
	return userGroups
}

// IsUserInGroup checks if a user is in a group
func (r *PostgresRepository) IsUserInGroup(userID, groupID string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user_groups WHERE user_id = $1 AND group_id = $2)`
	_ = r.db.QueryRow(query, userID, groupID).Scan(&exists)
	return exists
}

// GetMemberCount returns the number of members in a group
func (r *PostgresRepository) GetMemberCount(groupID string) int {
	var count int
	query := `SELECT COUNT(*) FROM user_groups WHERE group_id = $1`
	_ = r.db.QueryRow(query, groupID).Scan(&count)
	return count
}

// GetUserRole returns the role of a user in a group
func (r *PostgresRepository) GetUserRole(userID, groupID string) (string, error) {
	var role string
	query := `SELECT role FROM user_groups WHERE user_id = $1 AND group_id = $2`
	
	err := r.db.QueryRow(query, userID, groupID).Scan(&role)
	if err == sql.ErrNoRows {
		return "", errors.New("user not in group")
	}
	return role, err
}
