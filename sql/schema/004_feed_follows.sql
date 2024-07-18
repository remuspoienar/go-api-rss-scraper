-- +goose Up
create table feed_follows
(
    "feed_id"    uuid      not null references feeds (id) on delete cascade,
    "user_id"    uuid      not null references users (id) on delete cascade,
    "followed_at" timestamp not null,

    PRIMARY KEY (feed_id, user_id)
);
-- +goose Down
drop table feed_follows;