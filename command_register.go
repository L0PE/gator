package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/L0PE/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Not enouf argments provided")
	}

	username := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if user != (database.User{}) {
		return fmt.Errorf("User already exist")
	}

	user, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		Name: username,
	})
	if err != nil {
		return err	
	}
	
	s.conf.SetUser(user.Name)

	fmt.Printf("User %s was created\n", username)
	fmt.Printf("User: %v\n", user)
 
	return nil
}
