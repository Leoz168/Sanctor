package community

import "errors"

var (
	ErrNotMember        = errors.New("user is not a member of this community")
	ErrAlreadyMember    = errors.New("user is already a member of this community")
	ErrCommunityNotFound = errors.New("community not found")
	ErrUnauthorized     = errors.New("user does not have permission to perform this action")
	ErrInvalidRole      = errors.New("invalid role: must be member, admin, or owner")
	ErrOwnerCannotLeave = errors.New("owner cannot leave community with other members")
)
