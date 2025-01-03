package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command, ctx context.Context) error {
	if err := s.db.DeleteAllUser(ctx); err != nil {
		return fmt.Errorf("Cannot reset user ,%v", err)
	}
	if err := s.db.DeleteAllFeed(ctx); err != nil {
		return fmt.Errorf("Cannot reset feed ,%v", err)
	}
	if err := s.db.DeleteAllFollow(ctx); err != nil {
		return fmt.Errorf("Cannot reset follow ,%v", err)
	}
	if err := s.db.DeleteAllPosts(ctx); err != nil {
		return fmt.Errorf("Cannot reset follow ,%v", err)
	}
	fmt.Println("Reset succes full.")
	return nil
}
