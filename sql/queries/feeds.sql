-- name: CreateFeed :one
insert into feeds(id, created_at, updated_at, url, name, user_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
select *
from feeds;

-- name: FollowFeed :one
insert into feed_follows(feed_id, user_id, followed_at)
values ($1, $2, $3)
returning *;

-- name: UnfollowFeed :exec
delete
from feed_follows
where feed_id = $1
  AND user_id = $2;

-- name: GetFeedFollow :one
select *
from feed_follows
where feed_id = $1
  AND user_id = $2;