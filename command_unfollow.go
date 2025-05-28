package main

import (
	"context"
	"fmt"

	"github.com/L0PE/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Not enough arguments")
	}

	url := cmd.args[0]
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url: url,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s user follow to feed %s deleted\n", user.Name, url)

	return nil
}
