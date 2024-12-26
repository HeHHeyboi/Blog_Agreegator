-- name: CreateFeed :one
INSERT INTO feeds (id,name,url,user_id)
VALUES (
	$1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeleteAllFeed :exec
DELETE FROM feeds;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE feeds.url = $1 LIMIT 1;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE feeds.id = $1 LIMIT 1;

-- name: GetFeeds :many
SELECT * FROM feeds;
