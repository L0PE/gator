-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
	INSERT INTO feed_follows (
		id,
		created_at,
		updated_at,
		user_id,
		feed_id
	) VALUES ($1, $2, $3, $4, $5)
		RETURNING *
) SELECT 
	inserted_feed_follow.*,
	users.name as user_name,
	feeds.name as feed_name
FROM inserted_feed_follow
INNER JOIN users
ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds
ON inserted_feed_follow.feed_id = feeds.id;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
	WHERE feed_follows.user_id = $1 
	AND feed_follows.feed_id = (
		SELECT id 
		FROM feeds 
		WHERE url = $2 
		LIMIT 1
	);
