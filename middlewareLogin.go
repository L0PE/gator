package main

import (
	"context"

	"github.com/L0PE/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state,cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.conf.Current_user_name);
		if err != nil {
			return err
		}

		return 	handler(s, cmd, currentUser)
	}
}
