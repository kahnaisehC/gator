package main

import (
	"context"
	"fmt"

	"github.com/kahnaisehC/blog_aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	followedFeeds, err := s.db.GetFeedFollowsByUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, f := range followedFeeds {
		fmt.Println(f.Name_2)
	}

	return nil
}
