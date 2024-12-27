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

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW()
WHERE feeds.id = $1;

-- name: MarkFollowFetched :exec
UPDATE feed_follows
SET updated_at = NOW()
WHERE feed_id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;


