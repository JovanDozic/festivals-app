package models

import "errors"

var (
	ErrUsernameEmpty     = errors.New("username cannot be empty")
	ErrPasswordEmpty     = errors.New("password hash cannot be empty")
	ErrRoleIDEmpty       = errors.New("role ID cannot be empty")
	ErrNameEmpty         = errors.New("name cannot be empty")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateUser     = errors.New("user already exists")
	ErrRoleNotFound      = errors.New("role does not exist")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrEmptyUsername     = errors.New("username cannot be empty")
	ErrEmptyPassword     = errors.New("password cannot be empty")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrCompanyCountZero  = errors.New("company count is zero")
	ErrBadRequest        = errors.New("bad request")
	ErrNotFound          = errors.New("not found")
	ErrPageNotFound      = errors.New("404 page not found")
	ErrInternal          = errors.New("internal error")
	ErrUserNotValid      = errors.New("user is not valid")
)
