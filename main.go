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
		list: make(map[string]func(*state, command) error),
	}	

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerList)
	cmds.register("agg", handlerAggragaion)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFields)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", handlerFollowings)

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
