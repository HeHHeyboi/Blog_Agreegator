package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"blog_agreegator/internal/database"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command, ctx context.Context) error {
	if len(cmd.arg) < 1 {
		return fmt.Errorf("Username is require")
	}
	name := cmd.arg[0]
	exist_user, _ := s.db.GetUser(ctx, name)
	if exist_user.Name != "" {
		return fmt.Errorf("Username %s already exist", name)
	}
	newUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return err
	}
	log.Print(newUser)
	err = s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("Cannot config username")
	}

	return nil
}
