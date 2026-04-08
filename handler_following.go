package main

import (
	"context"
	"fmt"

	"github.com/Wayne_Francis/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Following takes no arguments\n")
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("*%v\n", feed.FeedName)
	}
	return nil
}
