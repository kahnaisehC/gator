-- name: CreateFeed :one
INSERT INTO feeds(url, name, user_id) VALUES(
	$1,
	$2,
	$3
)
RETURNING *;

-- name: GetFeeds :many
SELECT *
FROM feeds;
