package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Wayne_Francis/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) > 0 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = int32(n)
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}
	for _, post := range posts {
		title := ""
		if post.Title.Valid {
			title = post.Title.String
		}
		desc := ""
		if post.Description.Valid {
			desc = post.Description.String
		}

		fmt.Printf(" - %s\n   %s\n   %s\n\n", title, post.Url, desc)
	}
	return nil
}
