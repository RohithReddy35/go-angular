package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

// UserRepository is a struct that holds the database connection
type UserRepository struct {
	DB *sql.DB
	QueryBuilder sq.StatementBuilderType
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
		QueryBuilder: sq.StatementBuilderType(sq.StatementBuilder.PlaceholderFormat(sq.Dollar)),
	}
}