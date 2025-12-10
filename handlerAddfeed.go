package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kahnaisehC/gator/internal/database"
	_ "github.com/lib/pq"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return errors.New("not enough arguments for add feed command. Need two: <name> <url>")
	}

	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]

	createFeedParams := database.CreateFeedParams{
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	}

	newFeed, err := s.db.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return err
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    newFeed.FeedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println("the feed is:")
	fmt.Println(newFeed)

	return nil
}
