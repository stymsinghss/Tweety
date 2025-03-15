package utils

import "errors"

// Common error messages used throughout the app
var (
	// ErrUserNotFound -> when user is not found in db
	ErrUserNotFound = errors.New("user not found")

	// ErrInvalidEmail -> when email fails regex check
	ErrInvalidEmail = errors.New("invalid email")

	// ErrInvalidUsername -> when username fails regex check
	ErrInvalidUsername = errors.New("invalid username")

	// ErrEmailTaken -> when email already exists
	ErrEmailTaken = errors.New("email already exists")

	// ErrUsernameTaken -> when username already exists
	ErrUsernameTaken = errors.New("username already exists")
)
