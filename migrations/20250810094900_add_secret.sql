-- +goose Up
-- +goose StatementBegin
CREATE TABLE secrets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    archived BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    data_type SMALLINT NOT NULL,
    content BYTEA NOT NULL,
    metadata JSONB
);

CREATE INDEX idx_secrets_user_id ON secrets(user_id);
CREATE INDEX idx_secrets_archived ON secrets(archived);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE secrets;
-- +goose StatementEnd
