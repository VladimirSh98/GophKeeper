package user

import (
	"context"
	"database/sql"
)

// Repo struct
type Repo struct {
	Conn *sql.DB
}

// Repository interface
type Repository interface {
	GetUserByLogin(ctx context.Context, login string, archived bool) (User, error)
	Create(ctx context.Context, login string, password string) (int, error)
	Delete(ctx context.Context, login string) error
}

// NewRepository create new repository
func NewRepository(conn *sql.DB) Repository {
	return &Repo{Conn: conn}
}
