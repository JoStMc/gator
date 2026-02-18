-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users WHERE name = $1;


-- name: ListUsers :many
SELECT name FROM users;

-- name: GetFeedFollowsForUser :many
SELECT 
    users.name AS userName,
    feeds.name AS feedName,
    feeds.url
FROM users
INNER JOIN feed_follows ON users.id = feed_follows.user_id
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE users.name = $1;
