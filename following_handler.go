package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, ctx context.Context, cmd command) error {
	user, err := s.db.GetUser(ctx, s.config.Username)
	if err != nil {
		return fmt.Errorf("Error Get User: %v", err)
	}

	following, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("Get following Error: %v", err)
	}
	fmt.Printf("%v's Following:\n", user.Name)
	for _, f := range following {
		feed, err := s.db.GetFeedByID(ctx, f.FeedID)
		if err != nil {
			return err
		}

		fmt.Println("  ", feed.Name)
	}
	return nil
}
