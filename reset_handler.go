package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	if err := s.db.DeleteAllUser(ctx); err != nil {
		return fmt.Errorf("Cannot reset user ,%v", err)
	}
	fmt.Println("Reset succes full.")
	return nil
}
