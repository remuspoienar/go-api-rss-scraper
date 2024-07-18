-- name: CreatePost :exec
insert into posts(id, created_at, updated_at, published_at, title, url, description, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetPostByUrl :one
select *
from posts
where url in ($1);

-- name: GetPostsForUser :many
select posts.id,
       posts.created_at,
       posts.updated_at,
       posts.published_at,
       posts.title,
       posts.url,
       posts.description,
       posts.feed_id
from feed_follows
         inner join feeds on feed_follows.feed_id = feeds.id
         inner join posts on feeds.id = posts.feed_id
where feed_follows.user_id = $1
limit $2;