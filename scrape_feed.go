package main

import (
	"blog_agreegator/internal/database"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func scrapeFeeds(s *state, ctx context.Context) (database.Feed, error) {
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return nextFeed, fmt.Errorf("Error Get feed: %v", err)
	}

	err = MarkFeed(s, ctx, nextFeed.ID)
	if err != nil {
		return nextFeed, err
	}

	rss, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return nextFeed, fmt.Errorf("Error Fetch feed: %v", err)
	}
	fmt.Println("Feed Item:")
	for _, item := range rss.Channel.Item {
		fmt.Printf("- Title: %s\n", item.Title)
	}

	return nextFeed, nil
}

func MarkFeed(s *state, ctx context.Context, id uuid.UUID) error {
	err := s.db.MarkFeedFetched(ctx, id)
	if err != nil {
		return fmt.Errorf("Error mark feed: %v", err)
	}
	err = s.db.MarkFollowFetched(ctx, id)
	if err != nil {
		return fmt.Errorf("Error mark follow: %v", err)
	}
	return nil
}
