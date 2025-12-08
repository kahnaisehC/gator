-- name: CreateFeedFollow :many
WITH created_feed_follow AS (
	INSERT INTO feed_follows  (user_id, feed_id, created_at, updated_at) 
	VALUES(
		$1,
		$2,
		$3,
		$4
	)
	RETURNING *
) SELECT * 
FROM feed_follows 
INNER JOIN users on users.id = feed_follows.user_id
INNER JOIN feeds on feed_follows.feed_id = feeds.feed_id;


-- name: GetFeedFollowsByUser :many
SELECT *
FROM feed_follows
INNER JOIN users on users.id = feed_follows.user_id
INNER JOIN feeds on feed_follows.feed_id = feeds.feed_id
WHERE users.ID = $1;


-- name: DeleteFollow :one
DELETE FROM feed_follows
WHERE $1 = feed_follows.user_id and $2 = feed_follows.feed_id
RETURNING *;

