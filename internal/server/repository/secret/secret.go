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
	query := "SELECT * FROM \"secrets\" s JOIN \"user\" u on u.id = s.user_id WHERE u.login = $1 and s.archived = $2;"
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
	secretID int,
	content []byte,
	metadata []byte,
	archived bool,
) (Secret, error) {
	var secret Secret
	query := "UPDATE \"secrets\" SET content = $1, metadata = $2, archived = $3 WHERE id = $4  RETURNING id, user_id, archived, created_at, updated_at, data_type, content, metadata;"
	err := repo.Conn.QueryRowContext(ctx, query, content, metadata, archived, secretID).Scan(&secret.ID, &secret.UserID, &secret.Archived, &secret.CreatedAt, &secret.UpdatedAt, &secret.DataType, &secret.Content, &secret.Metadata)
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
