package user

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetUserByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	createdAt := time.Now()
	rows := sqlmock.NewRows([]string{"id", "created_at", "login", "hash", "archived"}).
		AddRow(1, createdAt, "testuser", "hash123", false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE login = $1 and archived = $2`)).
		WithArgs("testuser", false).
		WillReturnRows(rows)

	u, err := repo.GetUserByLogin(context.Background(), "testuser", false)
	require.NoError(t, err)
	require.Equal(t, 1, u.ID)
	require.Equal(t, "testuser", u.Login)
	require.Equal(t, "hash123", u.Hash)
	require.False(t, u.Archived)
	require.WithinDuration(t, createdAt, u.CreatedAt, time.Second)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "user" (login, hash, archived) VALUES ($1, $2, false) RETURNING id;`)).
		WithArgs("newuser", "password123").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(42))

	id, err := repo.Create(context.Background(), "newuser", "password123")
	require.NoError(t, err)
	require.Equal(t, 42, id)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "user" SET archived = $1 WHERE login = $2`)).
		WithArgs(true, "olduser").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(context.Background(), "olduser")
	require.NoError(t, err)
}
