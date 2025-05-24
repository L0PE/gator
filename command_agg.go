package main

import (
	"context"
	"fmt"
)

const (
	url = "https://www.wagslane.dev/index.xml"
)

func handlerAggragaion(s *state, cmd command) error {
	feeds, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Printf("Feeds: %+v\n", feeds)

	return nil
}
