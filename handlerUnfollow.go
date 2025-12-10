package main

import (
	"context"
	"errors"

	"github.com/kahnaisehC/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("not enough arguments to unfollow, need a feed url")
	}

	feedUrl := cmd.arguments[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}
	unfollowParams := database.DeleteFollowParams{
		UserID: user.ID,
		FeedID: feed.FeedID,
	}
	_, err = s.db.DeleteFollow(context.Background(), unfollowParams)
	if err != nil {
		return err
	}

	return nil
}
