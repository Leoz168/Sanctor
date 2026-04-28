package dm

import (
	"database/sql"

	"sanctor/internal/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *database.DB) Repository {
	return &GormRepository{db: db.Gorm}
}

func (r *GormRepository) CreateGroup(group *DMGroup) error {
	return r.db.Create(group).Error
}

func (r *GormRepository) AddUserToGroup(groupUser *DMGroupUser) error {
	return r.db.Create(groupUser).Error
}

func (r *GormRepository) GetGroupUsers(groupID uuid.UUID) ([]*DMGroupUser, error) {
	var members []*DMGroupUser
	err := r.db.Where("group_id = ?", groupID).Order("joined_at asc").Find(&members).Error
	if err != nil {
		return nil, err
	}

	if len(members) == 0 {
		var count int64
		if err := r.db.Model(&DMGroup{}).Where("id = ?", groupID).Count(&count).Error; err != nil {
			return nil, err
		}
		if count == 0 {
			return nil, ErrDMGroupNotFound
		}
	}

	return members, nil
}

func (r *GormRepository) GetUserGroups(userID uuid.UUID) ([]*DMGroup, error) {
	var groups []*DMGroup
	err := r.db.
		Table("dm_groups").
		Joins("JOIN dm_group_users ON dm_group_users.group_id = dm_groups.id").
		Where("dm_group_users.user_id = ?", userID).
		Order("dm_groups.updated_at desc").
		Find(&groups).Error
	return groups, err
}

func (r *GormRepository) FindDirectGroupByUsers(userA, userB uuid.UUID) (*DMGroup, error) {
	query := `
		SELECT g.id, g.created_at, g.updated_at
		FROM dm_groups g
		JOIN dm_group_users gu ON gu.group_id = g.id
		WHERE gu.user_id IN (?, ?)
		GROUP BY g.id, g.created_at, g.updated_at
		HAVING COUNT(DISTINCT gu.user_id) = 2
		   AND (SELECT COUNT(*) FROM dm_group_users dgu WHERE dgu.group_id = g.id) = 2
		LIMIT 1
	`

	var group DMGroup
	err := r.db.Raw(query, userA, userB).Scan(&group).Error
	if err != nil {
		return nil, err
	}
	if group.ID == uuid.Nil {
		return nil, ErrDMGroupNotFound
	}

	return &group, nil
}

func (r *GormRepository) IsUserInGroup(userID, groupID uuid.UUID) bool {
	var count int64
	err := r.db.Model(&DMGroupUser{}).Where("user_id = ? AND group_id = ?", userID, groupID).Count(&count).Error
	return err == nil && count > 0
}

func (r *GormRepository) SaveMessage(message *DMMessage) error {
	return r.db.Create(message).Error
}

func (r *GormRepository) GetGroupMessages(groupID uuid.UUID, limit int) ([]*DMMessage, error) {
	var messages []*DMMessage
	query := r.db.Where("group_id = ?", groupID).Order("message_time asc")
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.Find(&messages).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return []*DMMessage{}, nil
		}
		return nil, err
	}
	return messages, nil
}
