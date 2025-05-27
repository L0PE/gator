package main

import (
	"context"
	"fmt"
)

func handlerFollowings(s *state, cmd command) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), s.conf.Current_user_name)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(" - " + feedFollow.FeedName)
	}

	return nil
}
