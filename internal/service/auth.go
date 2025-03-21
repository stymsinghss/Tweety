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

type key string
const (
	// TokenLifespan -> 14 days token expiry
	TokenLifespan = time.Hour * 24 * 14

	// KeyAuthUserId to use in context
	KeyAuthUserId key = "auth_user_id"
)

type LoginOutput struct {
	Token     string
	ExpiresAt time.Time
	AuthUser  User
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
	s.token.SetTTL(uint32(TokenLifespan.Seconds()))
	if err != nil {
		return out, fmt.Errorf("could not create token for the user. Failed with -> %v\n", err)
	}

	// attach token expiry
	out.ExpiresAt = time.Now().Add(TokenLifespan)
	return out, nil
}

// AuthUserId -> decoded the token and returns the userId from the token
func (s *Service)  AuthUserId(token string) (int64, error) {
	str, err := s.token.DecodeToString(token)
	if err != nil {
		return 0, fmt.Errorf("could not decode token :%v/n", err)
	}

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse userId from the token :%v/n", err)
	}
	return i, nil
}


// temporary endpoint -> to get user details from the token

func (s *Service) AuthUser(ctx context.Context) (User, error) {
	var u User
	uid, ok := ctx.Value(KeyAuthUserId).(int64)
	if !ok {
		return u, utils.ErrUnauthenticated
	}

	query := "SELECT username FROM users where id = $1"
	err := s.db.QueryRowContext(ctx, query, uid).Scan(&u.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return u, utils.ErrUserNotFound
	}

	if err != nil {
		return u, fmt.Errorf("could not query user. Failed with -> %v\n", err)
	}

	u.ID = uid
	return u, nil
}