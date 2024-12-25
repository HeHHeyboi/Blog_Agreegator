-- name: CreateFeed :one
INSERT INTO feed (name,url,user_id)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: DeleteAllFeed :exec
DELETE FROM feed;
