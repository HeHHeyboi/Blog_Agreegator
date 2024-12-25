package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	ctx := context.Background()
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
