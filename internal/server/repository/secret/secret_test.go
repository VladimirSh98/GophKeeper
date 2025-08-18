package secret

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VladimirSh98/GophKeeper/internal/client/repository/secret"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

func TestCreateSecret(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	content := []byte("test content")
	metadata := []byte(`{"key":"value"}`)

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "secrets" (user_id, data_type, content, metadata, archived) VALUES ($1, $2, $3, $4, false) RETURNING id, user_id, archived, created_at, updated_at, data_type, content, metadata;`)).
		WithArgs(1, secret.LoginPassword, content, metadata).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "archived", "created_at", "updated_at", "data_type", "content", "metadata"}).
			AddRow(42, 1, false, time.Now(), time.Now(), secret.LoginPassword, content, metadata))

	s, err := repo.Create(context.Background(), 1, int(secret.LoginPassword), content, metadata)
	require.NoError(t, err)
	require.Equal(t, 42, s.ID)
	require.Equal(t, 1, s.UserID)
	require.Equal(t, content, s.Content)
}

func TestGetSecretsByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	content := []byte("test content")
	metadata := []byte(`{"key":"value"}`)

	rows := sqlmock.NewRows([]string{"id", "user_id", "archived", "created_at", "updated_at", "data_type", "content", "metadata"}).
		AddRow(1, 1, false, time.Now(), time.Now(), secret.LoginPassword, content, metadata)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT s.* FROM "secrets" s JOIN "user" u on u.id = s.user_id WHERE u.login = $1 and s.archived = $2;`)).
		WithArgs("testuser", false).
		WillReturnRows(rows)

	secrets, err := repo.GetSecretsByUser(context.Background(), "testuser")
	require.NoError(t, err)
	require.Len(t, secrets, 1)
	require.Equal(t, content, secrets[0].Content)
}

func TestUpdateByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	content := []byte("updated content")
	metadata := []byte(`{"key":"updated"}`)

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "secrets" s SET content = $1, metadata = $2 FROM "user" u WHERE s.id = $3 AND u.login = $4 AND s.user_id = u.id RETURNING s.id, s.user_id, s.archived, s.created_at, s.updated_at, s.data_type, s.content, s.metadata;`)).
		WithArgs(content, metadata, 1, "testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "archived", "created_at", "updated_at", "data_type", "content", "metadata"}).
			AddRow(1, 1, false, time.Now(), time.Now(), secret.LoginPassword, content, metadata))

	s, err := repo.UpdateByID(context.Background(), "testuser", 1, content, metadata)
	require.NoError(t, err)
	require.Equal(t, content, s.Content)
}

func TestDeleteByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "secrets" s SET archived = $1 FROM "user" u  WHERE u.id = s.user_id AND u.login = $2`)).
		WithArgs(true, "testuser").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteByLogin(context.Background(), "testuser")
	require.NoError(t, err)
}

func TestDeleteByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	content := []byte("test content")
	metadata := []byte(`{"key":"value"}`)

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "secrets" s SET archived = true FROM "user" u WHERE s.id = $1 AND u.login = $2 AND s.user_id = u.id RETURNING s.id, s.user_id, s.archived, s.created_at, s.updated_at, s.data_type, s.content, s.metadata;`)).
		WithArgs(1, "testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "archived", "created_at", "updated_at", "data_type", "content", "metadata"}).
			AddRow(1, 1, true, time.Now(), time.Now(), secret.LoginPassword, content, metadata))

	s, err := repo.DeleteByID(context.Background(), "testuser", 1)
	require.NoError(t, err)
	require.True(t, s.Archived)
}

func TestGetSecretByIDUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &Repo{Conn: db}

	content := []byte("test content")
	metadata := []byte(`{"key":"value"}`)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT s.* FROM "secrets" s JOIN "user" u on u.id = s.user_id WHERE u.login = $1 and s.id = $2 and s.archived = $3;`)).
		WithArgs("testuser", 1, false).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "archived", "created_at", "updated_at", "data_type", "content", "metadata"}).
			AddRow(1, 1, false, time.Now(), time.Now(), secret.LoginPassword, content, metadata))

	s, err := repo.GetSecretByIDUser(context.Background(), 1, "testuser")
	require.NoError(t, err)
	require.Equal(t, content, s.Content)
}
