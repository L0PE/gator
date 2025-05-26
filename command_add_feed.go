package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/L0PE/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Not enough arguments")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.conf.Current_user_name);
	if err != nil {
		return err
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name: name,
		Url: url,
		UserID: currentUser.ID,
	})
	if err != nil {
		return err
	}

    fmt.Println("Feed added:")
    fmt.Println(" - Name:", feed.Name)
    fmt.Println(" - URL:", feed.Url)
    fmt.Println(" - User ID:", feed.UserID)
    fmt.Println(" - Created At:", feed.CreatedAt.Time)
    fmt.Println(" - Updated At:", feed.UpdatedAt.Time)

	return nil
}
