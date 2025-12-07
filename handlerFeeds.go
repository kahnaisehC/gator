package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, f := range feeds {
		creator := ""
		user, err := s.db.GetUserByUUID(context.Background(), f.UserID)
		if err != nil {
			creator = "ERROR: creator not found"
		}
		creator = user.Name
		fmt.Println("Name: " + f.Name)
		fmt.Println("URL: " + f.Url)
		fmt.Println("Created by: " + creator)
	}
	return nil
}
