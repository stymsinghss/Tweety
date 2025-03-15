package service

import (
	"database/sql"
	"github.com/hako/branca"
)

// Service -> contains business logic
type Service struct {
	db    *sql.DB
	token *branca.Branca
}

// New -> returns an instance of Service
func New(db *sql.DB, token *branca.Branca) *Service {
	return &Service{
		db,
		token,
	}
}
