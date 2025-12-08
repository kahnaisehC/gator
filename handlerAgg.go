package main

import (
	"context"
	"database/sql"
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

	s.feeds[markedFeedRes.Url] = v

	for _, f := range s.feeds {
		if f == nil {
			continue
		}
		println(f.Channel.Title)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
