package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	// ErrUserNotFound -> when user is not found in db
	ErrUserNotFound = errors.New("user not found")
	
	// rxEmail -> regular expression for email validation
	rxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")

	// ErrInvalidEmail -> when email fails regex check
	ErrInvalidEmail = errors.New("invalid email")

	// rxUsername -> regular expression for username validation
	rxUsername = regexp.MustCompile("^[a-zA-Z][[a-zA-Z0-9_-]{0,17}$")

	// ErrInvalidUsername -> when username fails regex check
	ErrInvalidUsername = errors.New("invalid username")

	// ErrEmailTaken -> when email already exists
	ErrEmailTaken = errors.New("email already exists")

	// ErrUsernameTaken -> when username already exists
	ErrUsernameTaken = errors.New("username already exists")
)

// User -> represents User
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}

// CreateUser -> inserts a user in database
func (s *Service) CreateUser(ctx context.Context, email, username string) error {
	// Validate email
	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return ErrInvalidEmail
	}

	// Validate username
	username = strings.TrimSpace(username)
	if !rxUsername.MatchString(username) {
		return ErrInvalidUsername
	}

	// query
	query := "INSERT INTO users (email, username) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, email, username)

	unique := IsUniqueViolation(err)
	if unique && strings.Contains(err.Error(), "email") {
		return ErrEmailTaken
	}

	if unique && strings.Contains(err.Error(), "username") {
		return ErrUsernameTaken
	}

	if err != nil {
		return fmt.Errorf("could not insert user: %v\n", err)
	}
	return nil
}