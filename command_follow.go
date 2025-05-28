package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/L0PE/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Not enough arguments")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err	
	}

    fmt.Println("Feed follow added:")
    fmt.Println(" - ID:", feedFollow.ID)
    fmt.Println(" - Created At:", feedFollow.CreatedAt.Time)
    fmt.Println(" - Updated At:", feedFollow.UpdatedAt.Time)
    fmt.Println(" - User Name:", feedFollow.UserName)
    fmt.Println(" - Feed Name:", feedFollow.FeedName)
	return nil
}
