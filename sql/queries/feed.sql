-- name: CreateFeed :one
INSERT INTO feeds(url, name, user_id) VALUES(
	$1,
	$2,
	$3
)
RETURNING *;

-- name: GetFeeds :many
SELECT *
FROM feeds
INNER JOIN users on feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1
LIMIT 1;

-- name: MarkFeedFetched :one
UPDATE  feeds
SET last_fetched_at = $1
WHERE feeds.feed_id = $2
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
