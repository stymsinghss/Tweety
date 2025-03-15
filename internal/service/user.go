package service

import (
	"context"
	"fmt"
	"github.com/stymsinghss/Tweety/internal/utils"
	"strings"
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
	if !utils.RxEmail.MatchString(email) {
		return utils.ErrInvalidEmail
	}

	// Validate username
	username = strings.TrimSpace(username)
	if !utils.RxUsername.MatchString(username) {
		return utils.ErrInvalidUsername
	}

	// query
	query := "INSERT INTO users (email, username) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, email, username)

	unique := IsUniqueViolation(err)
	if unique && strings.Contains(err.Error(), "email") {
		return utils.ErrEmailTaken
	}

	if unique && strings.Contains(err.Error(), "username") {
		return utils.ErrUsernameTaken
	}

	if err != nil {
		return fmt.Errorf("could not insert user: %v\n", err)
	}
	return nil
}