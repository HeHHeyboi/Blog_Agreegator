// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feed (name,url,user_id)
VALUES (
    $1,
    $2,
    $3
)
RETURNING name, url, user_id
`

type CreateFeedParams struct {
	Name   string
	Url    string
	UserID uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed, arg.Name, arg.Url, arg.UserID)
	var i Feed
	err := row.Scan(&i.Name, &i.Url, &i.UserID)
	return i, err
}

const deleteAllFeed = `-- name: DeleteAllFeed :exec
DELETE FROM feed
`

func (q *Queries) DeleteAllFeed(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllFeed)
	return err
}
