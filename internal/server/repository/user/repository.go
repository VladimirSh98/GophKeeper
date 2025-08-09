package user

import (
	"context"
	"database/sql"
)

type Repo struct {
	Conn *sql.DB
}

type Repository interface {
	GetUserByLogin(ctx context.Context, login string, archived bool) (User, error)
	Create(ctx context.Context, login string, password string) (int, error)
}

func NewRepository(conn *sql.DB) Repository {
	return &Repo{Conn: conn}
}
