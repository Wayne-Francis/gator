package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("please enter time between feeds\n")
	}
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %+v\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}

func scrapeFeeds(s *state) {
	next_feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Printf("Fetching feed: %v\n", next_feed.Name)
	_, err = s.db.MarkFeedFetched(context.Background(), next_feed.ID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	feed, err := fetchFeed(context.Background(), next_feed.Url)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	for i := range feed.Channel.Item {
		fmt.Printf("*%v\n", feed.Channel.Item[i].Title)
	}

}
