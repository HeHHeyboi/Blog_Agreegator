package main

import (
	"blog_agreegator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddfeed(s *state, ctx context.Context, cmd command, user database.User) error {
	if len(cmd.arg) < 2 {
		return fmt.Errorf("Please input name & url link")
	}
	name := cmd.arg[0]
	url := cmd.arg[1]

	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("Error when create feed: %v", err)
	}
	_, err = s.db.CreateFeedFollow(
		ctx,
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    newFeed.ID,
		},
	)

	return nil
}
