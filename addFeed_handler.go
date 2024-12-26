package main

import (
	"blog_agreegator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddfeed(name, url string, s *state, ctx context.Context) (database.Feed, error) {
	username := s.config.Username
	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return database.Feed{}, fmt.Errorf("GetUser error: %v", err)
	}
	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return database.Feed{}, fmt.Errorf("Error when create feed: %v", err)
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

	return newFeed, nil
}
