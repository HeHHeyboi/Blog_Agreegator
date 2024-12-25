package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerFeeds(s *state, ctx context.Context) error {
	feed_list, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}
	for _, f := range feed_list {
		user, err := getUserID(s, ctx, f.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Title: %s\nURL: %s\nUser: %s\n\n", f.Name, f.Url, user)
	}
	return nil
}

func getUserID(s *state, ctx context.Context, id uuid.UUID) (string, error) {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return "", err
	}
	return user.Name, nil

}
