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
		fmt.Println("Name: " + f.Name)
		fmt.Println("URL: " + f.Url)
		fmt.Println("Created by: " + f.Name_2)
	}
	return nil
}
