package main

import (
	"context"
	"fmt"

	"github.com/L0PE/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Not enouf argments provided")
	}

	username := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), username)
	if user == (database.User{}){
		return fmt.Errorf("User not found")
	}


	err = s.conf.SetUser(username)
	if err != nil {
		return err	
	}

	fmt.Printf("Current user set to %s\n", username)

	return nil
}
