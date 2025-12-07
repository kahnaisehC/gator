package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kahnaisehC/blog_aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("the register command expects a name as an argument")
	}
	name := cmd.arguments[0]

	userParams := database.CreateUserParams{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        uuid.New(),
	}

	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	s.cfg.SetUser(name)
	fmt.Println(name + " Has been registered successfully!!")

	return nil
}
