package service

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

// IsUniqueViolation checks if the given error is a PostgreSQL unique violation error.
func IsUniqueViolation(err error) bool {
	var pgerr *pgconn.PgError
	ok := errors.As(err, &pgerr)
	return ok && pgerr.Code == "23505"
}
