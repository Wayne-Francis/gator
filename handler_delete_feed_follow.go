package main

import (
	"context"
	"fmt"

	"github.com/Wayne_Francis/gator/internal/database"
)

func handlerDeleteFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Please enter a feed to follow\n")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
