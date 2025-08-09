package user

import (
	"context"
)

// GetUserByLogin get user by login
func (repo *Repo) GetUserByLogin(ctx context.Context, login string, archived bool) (User, error) {
	var record User
	query := "SELECT * FROM \"user\" WHERE login = $1 and archived = $2"
	row := repo.Conn.QueryRowContext(ctx, query, login, archived)
	err := row.Scan(&record.ID, &record.CreatedAt, &record.Login, &record.Hash, &record.Archived)
	if err != nil {
		return record, err
	}
	return record, nil
}

// Create new user
func (repo *Repo) Create(ctx context.Context, login string, password string) (int, error) {
	query := "INSERT INTO \"user\" (login, hash, archived) VALUES ($1, $2, false) RETURNING id;"
	var ID int
	err := repo.Conn.QueryRowContext(ctx, query, login, password).Scan(&ID)

	if err != nil {
		return 0, err
	}
	return ID, nil
}

// Delete user
func (repo *Repo) Delete(ctx context.Context, login string) error {
	query := "UPDATE \"user\" SET archived = true WHERE login = $1"
	_, err := repo.Conn.ExecContext(ctx, query, login)
	if err != nil {
		return err
	}
	return nil
}
