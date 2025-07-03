package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/L0PE/gator/internal/database"
	"github.com/google/uuid"
)

const (
	url = "https://www.wagslane.dev/index.xml"
)

func handlerAggragaion(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Not enough arguments")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs.String())
	ticker := time.NewTicker(time_between_reqs)

	for ; ; <- ticker.C {
		scrapeFeeds(s)			
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return nil
	}

	fmt.Printf("Fetching %s", nextFeed.Url)
	feeds, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err	
	}

	fmt.Printf("%s fetched\n", feeds.Channel.Title)
	for _, feed := range feeds.Channel.Items {
		post, err := s.db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID: uuid.New(),
				CreatedAt: sql.NullTime{
					Time: time.Now().UTC(),
					Valid: true,
				},
			UpdatedAt: sql.NullTime{
				Time: time.Now().UTC(),
				Valid: true,
			},
			Title: feed.Title,
			Url: feed.Link,
			Description: feed.Description,
			PublishedAt: time.Now().UTC(),
			FeedID: nextFeed.ID,
		},
	)
	if err != nil {
		fmt.Printf("Error while saving post: %v", err)
		continue
	}

	fmt.Printf(" - Title: %s", post.Title)
	}

	return nil
}
