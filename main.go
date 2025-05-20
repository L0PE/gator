package main

import (
	"fmt"
	"os"

	"github.com/L0PE/gator/internal/config"
)


func main() {
	configStruct, err := config.Read()
	if err != nil {
		fmt.Printf("Errow durind reading the config: %v", err)
		return
	}
	
	s := state {
		&configStruct,
	}

	cmds := commands{
		map[string]func(*state, command) error {
			"login": handlerLogin,
		},
	}
	
	if len(os.Args) < 2 {
		fmt.Printf("Please provide at least two arguments\n")
		os.Exit(1)	
	}

	command := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(&s, command)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
