-- name: CreatePost :one
INSERT INTO posts (created_at, updated_at, title, url, description, published_at, feed_id) VALUES(
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
)
RETURNING *;

-- name: GetPostsForUsers :many
SELECT DISTINCT *
FROM posts
INNER JOIN feeds on feeds.feed_id = posts.feed_id
INNER JOIN feed_follows on feed_follows.feed_id = feeds.feed_id
INNER JOIN users on feed_follows.user_id = users.id
WHERE users.id = $1 
ORDER BY posts.published_at ASC
LIMIT $2;

