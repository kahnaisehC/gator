package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/kahnaisehC/blog_aggregator/internal/database"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	markFeedParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
		FeedID: nextFeed.FeedID,
	}

	markedFeedRes, err := s.db.MarkFeedFetched(context.Background(), markFeedParams)
	if err != nil {
		return err
	}

	v, err := fetchFeed(context.Background(), markedFeedRes.Url)
	if err != nil {
		return err
	}

	fmt.Println(v.Channel.Title)
	for _, item := range v.Channel.Item {
		fmt.Println("- " + item.Title)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return errors.New("not enough arguments for agg command: need a time between requests")
	}
	timeString := cmd.arguments[0]
	timeBetweenReqs, err := time.ParseDuration(timeString)
	if err != nil {
		fmt.Println("invalid time between requests: defaulted to 5s per request")
		timeBetweenReqs = time.Second * 5
	}
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
