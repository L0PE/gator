package main

import (
	"context"
	"fmt"
)

func handlerList(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil{
		return err
	}

	for _, user := range users {
		result := fmt.Sprintf(" * %s", user.Name)
		if user.Name == s.conf.Current_user_name {
			result += " (current)"
		}

		fmt.Println(result)
	}

	return nil
}
