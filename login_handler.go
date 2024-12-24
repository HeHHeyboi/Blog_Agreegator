package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arg) < 1 {
		return fmt.Errorf("Username is require")
	}

	name := cmd.arg[0]
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, name)
	if err != nil {
		return fmt.Errorf("User doesn't exit")
	}
	err = s.config.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
