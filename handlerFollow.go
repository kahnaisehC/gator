package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kahnaisehC/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("not enough arguments for the follow handler: needs an url")
	}

	url := cmd.arguments[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	createFeedParams := database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.FeedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.db.CreateFeedFollow(context.Background(), createFeedParams)
	if err != nil {
		return err
	}
	feed, err = s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Println(s.cfg.CurrentUserName + " now follows " + feed.Name)

	return nil
}
