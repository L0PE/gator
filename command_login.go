package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Not enouf argments provided")
	}

	username := cmd.args[0]

	err := s.conf.SetUser(username)
	if err != nil {
		return err	
	}

	fmt.Printf("Current user set to %s\n", username)

	return nil
}
