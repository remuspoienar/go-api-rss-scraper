-- +goose Up
CREATE TABLE users
(
    "id"         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name"       text NOT NULL,
    "created_at" timestamp WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    "updated_at" timestamp WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

-- +goose Down
DROP TABLE users;