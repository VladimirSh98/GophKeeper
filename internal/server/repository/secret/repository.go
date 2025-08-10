package secret

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
	Create(
		ctx context.Context,
		userID int,
		dataType int,
		content []byte,
		metadata []byte,
	) (Secret, error)
	GetSecretsByUser(ctx context.Context, login string) ([]Secret, error)
	UpdateByID(
		ctx context.Context,
		secretID int,
		content []byte,
		metadata []byte,
		archived bool,
	) (Secret, error)
	DeleteByLogin(ctx context.Context, login string) error
}

// NewRepository create new repository
func NewRepository(conn *sql.DB) Repository {
	return &Repo{Conn: conn}
}
