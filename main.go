package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/L0PE/gator/internal/config"
	"github.com/L0PE/gator/internal/database"

	_ "github.com/lib/pq"
)


func main() {
	configStruct, err := config.Read()
	if err != nil {
		fmt.Printf("Errow durind reading the config: %v", err)
		return
	}
	
	db, err := sql.Open("postgres", configStruct.DB_url)
	if err != nil {
		fmt.Printf("Could not connect to the database: %v\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	s := state {
		&configStruct,
		dbQueries,
	}

	cmds := commands{
		map[string]func(*state, command) error {
			"login": handlerLogin,
			"register": handlerRegister,
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
