package service

import "errors"

var (
	// ErrUserNotFound -> when user is not found in db
	ErrUserNotFound = errors.New("user not found")
	
)

// User -> represents User
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}