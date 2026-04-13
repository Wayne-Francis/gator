package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Wayne_Francis/gator/internal/database"
	"github.com/google/uuid"
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
		item := feed.Channel.Item[i]

		parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)

		var publishedAt sql.NullTime
		if err == nil {
			publishedAt = sql.NullTime{Time: parsedTime, Valid: true}
		} else {
			publishedAt = sql.NullTime{Valid: false}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid:  item.Title != "",
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: publishedAt,
			FeedID:      next_feed.ID,
		})
		if err != nil {
			msg := err.Error()
			if strings.Contains(msg, "posts_url_key") {
			} else {
				fmt.Println(err)
			}
		}
	}
}
