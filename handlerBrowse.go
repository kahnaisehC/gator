package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kahnaisehC/blog_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) > 1 {
		val, err := strconv.Atoi(cmd.arguments[0])
		if err == nil {
			limit = val
		}
	}
	params := database.GetPostsForUsersParams{
		ID:    user.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUsers(context.Background(), params)
	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Println("--- " + p.Title + " ---")
		if p.Description.Valid {
			fmt.Println("Description: " + p.Description.String)
		}
	}

	return nil
}
