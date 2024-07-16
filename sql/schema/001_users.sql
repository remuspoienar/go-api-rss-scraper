-- +goose Up
CREATE TABLE users
(
    "id"         uuid primary key,
    "name"       text                     not null,
    "created_at" timestamp with time zone not null default (current_timestamp at time zone 'UTC'),
    "updated_at" timestamp with time zone not null default (current_timestamp at time zone 'UTC')
);

-- +goose Down
DROP TABLE users;