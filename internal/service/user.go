package service

import (
	"errors"
	"regexp"
)

var (
	// ErrUserNotFound -> when user is not found in db
	ErrUserNotFound = errors.New("user not found")
	
	// rxEmail -> regular expression for email validation
	rxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")

	// ErrInvalidEmail -> when email fails regex check
	ErrInvalidEmail = errors.New("invalid email")
)

// User -> represents User
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}