-- +goose Up
CREATE TABLE feeds
(
    "id"         uuid primary key,
    "created_at" timestamp with time zone not null default (current_timestamp at time zone 'UTC'),
    "updated_at" timestamp with time zone not null default (current_timestamp at time zone 'UTC'),
    "url"        text unique              not null,
    "name"       text                     not null,
    "user_id"    uuid                     not null references users (id) on delete cascade

);

-- +goose Down
DROP TABLE feeds;

