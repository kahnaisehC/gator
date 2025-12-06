package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/kahnaisehC/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return errors.New("not enough arguments for add feed command. Need two: <name> <url>")
	}
	username := s.cfg.CurrentUserName
	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	createFeedParams := database.CreateFeedParams{
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	}

	newFeed, err := s.db.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return err
	}
	fmt.Println("the feed is:")
	fmt.Println(newFeed)

	return nil
}
