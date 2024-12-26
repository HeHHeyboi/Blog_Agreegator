package main

import (
	"blog_agreegator/internal/database"
	"context"
	"fmt"
)

func middlewareLogin(handler func(s *state, ctx context.Context, cmd command, user database.User) error) func(s *state, ctx context.Context, cmd command) error {
	return func(s *state, ctx context.Context, cmd command) error {
		user, err := s.db.GetUser(ctx, s.config.Username)
		if err != nil {
			fmt.Printf("Cannot get user : %v", err)
			return nil
		}
		return handler(s, ctx, cmd, user)
	}
}
