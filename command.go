package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c commands) run(s *state, cmd command) error {
	handler, ok := c.list[cmd.name]
	if !ok {
		return fmt.Errorf("Command not found")
	}

	return handler(s, cmd)
}

func (c commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}
