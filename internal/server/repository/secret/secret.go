package secret

import "context"

// Create new secret
func (repo *Repo) Create(
	ctx context.Context,
	userID int,
	dataType int,
	content []byte,
	metadata []byte,
) (Secret, error) {
	var secret Secret
	query := "INSERT INTO \"secrets\" (user_id, data_type, content, metadata, archived) VALUES ($1, $2, $3, $4, false) RETURNING id, user_id, archived, created_at, updated_at, data_type, content, metadata;"
	err := repo.Conn.QueryRowContext(ctx, query, userID, dataType, content, metadata).Scan(&secret.ID, &secret.UserID, &secret.Archived, &secret.CreatedAt, &secret.UpdatedAt, &secret.DataType, &secret.Content, &secret.Metadata)

	if err != nil {
		return secret, err
	}
	return secret, nil
}

// GetSecretsByUser get not archived secrets by user
func (repo *Repo) GetSecretsByUser(ctx context.Context, login string) ([]Secret, error) {
	secrets := make([]Secret, 0)
	query := "SELECT s.* FROM \"secrets\" s JOIN \"user\" u on u.id = s.user_id WHERE u.login = $1 and s.archived = $2;"
	rows, err := repo.Conn.QueryContext(ctx, query, login, false)
	if err != nil {
		return secrets, err
	}
	err = rows.Err()
	defer rows.Close()
	if err != nil {
		return secrets, err
	}
	for rows.Next() {
		var secret Secret
		err = rows.Scan(&secret.ID, &secret.UserID, &secret.Archived, &secret.CreatedAt, &secret.UpdatedAt, &secret.DataType, &secret.Content, &secret.Metadata)
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, secret)
	}
	return secrets, nil
}

// UpdateByID update secret by ID
func (repo *Repo) UpdateByID(
	ctx context.Context,
	login string,
	secretID int,
	content []byte,
	metadata []byte,
) (Secret, error) {
	var secret Secret
	query := "UPDATE \"secrets\" s SET content = $1, metadata = $2 FROM \"user\" u WHERE s.id = $3 AND u.login = $4 AND s.user_id = u.id RETURNING s.id, s.user_id, s.archived, s.created_at, s.updated_at, s.data_type, s.content, s.metadata;"
	err := repo.Conn.QueryRowContext(ctx, query, content, metadata, secretID, login).Scan(&secret.ID, &secret.UserID, &secret.Archived, &secret.CreatedAt, &secret.UpdatedAt, &secret.DataType, &secret.Content, &secret.Metadata)
	if err != nil {
		return secret, err
	}
	return secret, nil
}

// DeleteByLogin delete secrets by login
func (repo *Repo) DeleteByLogin(
	ctx context.Context,
	login string,
) error {
	query := "UPDATE \"secrets\" s SET archived = $1 FROM \"user\" u  WHERE u.id = s.user_id AND u.login = $2"
	_, err := repo.Conn.ExecContext(ctx, query, true, login)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID delete secret by ID
func (repo *Repo) DeleteByID(
	ctx context.Context,
	login string,
	secretID int,
) (Secret, error) {
	var secret Secret
	query := "UPDATE \"secrets\" s SET archived = true FROM \"user\" u WHERE s.id = $1 AND u.login = $2 AND s.user_id = u.id RETURNING s.id, s.user_id, s.archived, s.created_at, s.updated_at, s.data_type, s.content, s.metadata;"
	err := repo.Conn.QueryRowContext(ctx, query, secretID, login).Scan(&secret.ID, &secret.UserID, &secret.Archived, &secret.CreatedAt, &secret.UpdatedAt, &secret.DataType, &secret.Content, &secret.Metadata)
	if err != nil {
		return secret, err
	}
	return secret, nil
}

// GetSecretByIDUser get not archived secret by id for user
func (repo *Repo) GetSecretByIDUser(ctx context.Context, secretID int, login string) (Secret, error) {
	query := "SELECT s.* FROM \"secrets\" s JOIN \"user\" u on u.id = s.user_id WHERE u.login = $1 and s.id = $2 and s.archived = $3;"
	row := repo.Conn.QueryRowContext(ctx, query, login, secretID, false)
	var secret Secret
	err := row.Scan(&secret.ID, &secret.UserID, &secret.Archived, &secret.CreatedAt, &secret.UpdatedAt, &secret.DataType, &secret.Content, &secret.Metadata)
	if err != nil {
		return secret, err
	}
	return secret, nil
}
