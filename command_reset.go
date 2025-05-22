package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil{
		return err
	}

	fmt.Printf("Users where deleted\n")

	return nil
}
