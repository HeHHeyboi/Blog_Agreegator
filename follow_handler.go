package main

import (
	"blog_agreegator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, ctx context.Context, cmd command, user database.User) error {
	if len(cmd.arg) < 1 {
		return fmt.Errorf("URL is required")
	}
	url := cmd.arg[0]

	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("URL Error: %v", err)
	}

	follow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Create Follow Error: %v", err)
	}

	fmt.Println("Follow Feed:", follow.FeedName)
	fmt.Println("by User:", follow.UserName)

	return nil
}
