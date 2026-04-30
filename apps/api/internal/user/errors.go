package user

import "errors"

var (
	ErrUserBlacklisted  = errors.New("user is blacklisted")
	ErrEmailBlacklisted = errors.New("email is blacklisted")
)
