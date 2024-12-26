package main

import (
	"blog_agreegator/internal/database"
	"context"
	"fmt"
)

func handlerUnfollow(s *state, ctx context.Context, cmd command, user database.User) error {
	if len(cmd.arg) < 1 {
		return fmt.Errorf("Please input url feed")
	}

	url := cmd.arg[0]
	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("Cannot find feed: %v", err)
	}

	err = s.db.DeleteFollow(ctx, database.DeleteFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error when delete feed: %v", err)
	}
	return nil
}
