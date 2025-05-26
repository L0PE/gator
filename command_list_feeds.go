package main

import (
	"context"
	"fmt"
)

func handlerListFields(s *state, _ command) error {
	feeds, err := s.db.GetFeedsWithUser(context.Background());
	if err != nil{
		return err
	}

	for _, feed := range feeds {
		result := fmt.Sprintf(" * name: %s url: %s user: %s", feed.Feed.Name, feed.Feed.Url, feed.User.Name)
		fmt.Println(result)
	}

	return nil
}
