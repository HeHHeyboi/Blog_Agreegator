package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command, ctx context.Context) error {
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		result := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.config.Username {
			result += " (current)"
		}
		fmt.Println(result)
	}
	return nil
}
