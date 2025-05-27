-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name) 
VALUES ( 
	$1,
	$2,
	$3,
	$4
) RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, name 
FROM users 
WHERE name = $1 
LIMIT 1;

-- name: ResetUsers :exec
TRUNCATE TABLE users CASCADE;

-- name: GetUsers :many
SELECT id, created_at, updated_at, name 
FROM users;

-- name: GetFeedFollowsForUser :many
SELECT 
	users.id as user_id, 
	users.name as user_name, 
	feeds.id as feed_id, 
	feeds.name as feed_name, 
	feeds.url as feed_url  
FROM users
LEFT JOIN feed_follows
ON users.id = feed_follows.user_id
INNER JOIN feeds
ON feeds.id = feed_follows.feed_id
WHERE users.name = $1;
