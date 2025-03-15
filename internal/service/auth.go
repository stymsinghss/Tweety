package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/stymsinghss/Tweety/internal/utils"
	"strconv"
	"strings"
	"time"
)

const (
	// TokenLifespan -> 14 days token expiry
	TokenLifespan = time.Hour * 24 * 14
)

type LoginOutput struct {
	Token string
	ExpiresAt time.Time
	AuthUser User
}

// Login -> logins the user by checking if the email is valid and attach tokens to it
func (s *Service) Login(ctx context.Context, email string) (LoginOutput, error) {
	var out LoginOutput

	// Validate email
	email = strings.TrimSpace(email)
	if !utils.RxEmail.MatchString(email) {
		return out, utils.ErrInvalidEmail
	}

	query := "SELECT id, username FROM users WHERE email = $1"
	err := s.db.QueryRowContext(ctx, query, email).Scan(&out.AuthUser.ID, &out.AuthUser.Username)

	// If no user found
	if errors.Is(err, sql.ErrNoRows) {
		return out, utils.ErrUserNotFound
	}

	if err != nil {
		return out, fmt.Errorf("could not query user: %v\n", err)
	}

	// attach tokens
	out.Token, err = s.token.EncodeToString(strconv.FormatInt(out.AuthUser.ID, 10))
	if err != nil {
		return out, fmt.Errorf("could not create token for the user. Failed with -> %v\n", err)
	}

	// attach token expiry
	out.ExpiresAt = time.Now().Add(TokenLifespan)
	return out, nil
}