package main

import (
	"blog_agreegator/internal/database"
	"context"
	"fmt"
)

func handlerAddfeed(name, url string, s *state, ctx context.Context) (database.Feed, error) {
	username := s.config.Username
	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return database.Feed{}, fmt.Errorf("GetUser error: %v", err)
	}
	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return database.Feed{}, fmt.Errorf("Error when create feed: %v", err)
	}

	return newFeed, nil
}
