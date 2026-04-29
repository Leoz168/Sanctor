package community

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

// NewPostgresRepository creates a new PostgreSQL community repository
func NewPostgresRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}

// Create creates a new community
func (r *PostgresRepository) Create(community *Community) error {
	query := `
		INSERT INTO communities (id, name, description, is_private, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query, community.ID, community.Name, community.Description, community.IsPrivate,
		community.CreatedBy, community.CreatedAt, community.UpdatedAt)
	return err
}

// FindByID finds a community by ID
func (r *PostgresRepository) FindByID(id uuid.UUID) (*Community, error) {
	community := &Community{}
	query := `SELECT id, name, description, is_private, created_by, created_at, updated_at 
	          FROM communities WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&community.ID, &community.Name, &community.Description,
		&community.IsPrivate, &community.CreatedBy, &community.CreatedAt, &community.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("community not found")
	}
	return community, err
}

// FindAll returns all communities
func (r *PostgresRepository) FindAll() []*Community {
	query := `SELECT id, name, description, is_private, created_by, created_at, updated_at 
	          FROM communities ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return []*Community{}
	}
	defer rows.Close()

	communities := []*Community{}
	for rows.Next() {
		community := &Community{}
		if err := rows.Scan(&community.ID, &community.Name, &community.Description, &community.IsPrivate,
			&community.CreatedBy, &community.CreatedAt, &community.UpdatedAt); err == nil {
			communities = append(communities, community)
		}
	}
	return communities
}

// Update updates an existing community
func (r *PostgresRepository) Update(community *Community) error {
	query := `UPDATE communities SET name = $2, description = $3, is_private = $4, updated_at = $5 
	          WHERE id = $1`

	result, err := r.db.Exec(query, community.ID, community.Name, community.Description,
		community.IsPrivate, community.UpdatedAt)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("community not found")
	}
	return nil
}

// Delete deletes a community (CASCADE will delete user_communities automatically)
func (r *PostgresRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM communities WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("community not found")
	}
	return nil
}

// AddUserToCommunity adds a user to a community
func (r *PostgresRepository) AddUserToCommunity(userCommunity *UserCommunity) error {
	query := `INSERT INTO user_communities (user_id, community_id, role, joined_at) 
	          VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(query, userCommunity.UserID, userCommunity.CommunityID,
		userCommunity.Role, userCommunity.JoinedAt)
	return err
}

// AddCommunityToInstitution links a community to an institution
func (r *PostgresRepository) AddCommunityToInstitution(communityInstitution *CommunityInstitution) error {
	query := `INSERT INTO community_institutions (community_id, institution_id, linked_at)
	          VALUES ($1, $2, $3)`

	_, err := r.db.Exec(query, communityInstitution.CommunityID, communityInstitution.InstitutionID, communityInstitution.LinkedAt)
	return err
}

// RemoveUserFromCommunity removes a user from a community
func (r *PostgresRepository) RemoveUserFromCommunity(userID, communityID uuid.UUID) error {
	query := `DELETE FROM user_communities WHERE user_id = $1 AND community_id = $2`
	result, err := r.db.Exec(query, userID, communityID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("user not in community")
	}
	return nil
}

// GetCommunityMembers returns all members of a community
func (r *PostgresRepository) GetCommunityMembers(communityID uuid.UUID) ([]*UserCommunity, error) {
	query := `SELECT user_id, community_id, role, joined_at 
	          FROM user_communities WHERE community_id = $1 ORDER BY joined_at`

	rows, err := r.db.Query(query, communityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := []*UserCommunity{}
	for rows.Next() {
		ug := &UserCommunity{}
		if err := rows.Scan(&ug.UserID, &ug.CommunityID, &ug.Role, &ug.JoinedAt); err == nil {
			members = append(members, ug)
		}
	}
	return members, nil
}

// GetUserCommunities returns all communities a user belongs to
func (r *PostgresRepository) GetUserCommunities(userID uuid.UUID) []*UserCommunity {
	query := `SELECT user_id, community_id, role, joined_at 
	          FROM user_communities WHERE user_id = $1 ORDER BY joined_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return []*UserCommunity{}
	}
	defer rows.Close()

	userCommunities := []*UserCommunity{}
	for rows.Next() {
		ug := &UserCommunity{}
		if err := rows.Scan(&ug.UserID, &ug.CommunityID, &ug.Role, &ug.JoinedAt); err == nil {
			userCommunities = append(userCommunities, ug)
		}
	}
	return userCommunities
}

// IsUserInCommunity checks if a user is in a community
func (r *PostgresRepository) IsUserInCommunity(userID, communityID uuid.UUID) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user_communities WHERE user_id = $1 AND community_id = $2)`
	_ = r.db.QueryRow(query, userID, communityID).Scan(&exists)
	return exists
}

// GetMemberCount returns the number of members in a community
func (r *PostgresRepository) GetMemberCount(communityID uuid.UUID) int {
	var count int
	query := `SELECT COUNT(*) FROM user_communities WHERE community_id = $1`
	_ = r.db.QueryRow(query, communityID).Scan(&count)
	return count
}

// GetUserRole returns the role of a user in a community
func (r *PostgresRepository) GetUserRole(userID, communityID uuid.UUID) (string, error) {
	var role string
	query := `SELECT role FROM user_communities WHERE user_id = $1 AND community_id = $2`

	err := r.db.QueryRow(query, userID, communityID).Scan(&role)
	if err == sql.ErrNoRows {
		return "", errors.New("user not in community")
	}
	return role, err
}
