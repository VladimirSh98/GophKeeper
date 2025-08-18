-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    created_at timestamp default NOW(),
    login text,
    hash text,
    archived bool
);
CREATE UNIQUE INDEX IF NOT EXISTS user_login_udx ON "user"(login) WHERE archived is false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX user_login_udx;
DROP TABLE "user";
-- +goose StatementEnd
