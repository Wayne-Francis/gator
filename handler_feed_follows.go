package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Wayne_Francis/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Please enter a URL to follow\n")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	follow_feed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed: Name: %v\nUser: %v\n", follow_feed.FeedName, follow_feed.UserName)
	return nil
}
